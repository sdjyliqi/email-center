import pymysql
import argparse
import chardet
import time
import datetime
from lxml import etree
import json
import os
import re
import eml_parser
#traverse_path...  读取目录下所有文件，生成列表
def traverse_path(path):
    list = []
    for fpathe,dirs,fs in os.walk(path):
          for f in fs:
            file_path = os.path.join(fpathe,f)
            list.append(file_path)
    return list


#filter_u_code...  unicode编码转换到正常中文
def filter_u_code(string):
    u_codes = re.findall(r'(\\u\w{4})',string)
    for u_code in u_codes:
        string = string.replace(u_code,u_code.encode("utf-8").decode("unicode_escape"))
    return string

#json_serial...   日期标准化
def json_serial(obj):
    if isinstance(obj, datetime.datetime):
        serial = obj.isoformat()
        return serial

#html_parser... 将从邮件正文的html中提取内容
def html_parser (original_content):
    selector=etree.HTML(original_content)
    content_list = selector.xpath('string(.)')
    content = ''
    for value in content_list:
        content += value.strip(' ')
    return content

#利用argparse包，从控制台传入参数
parser = argparse.ArgumentParser(description = '命令行中输入邮件解析所需参数')
parser.add_argument('--path', type = str, required = True, help = '邮件所在目录(必填)')
parser.add_argument('--manual', type = int, default = 0, help = '是否为正常邮件(可选)，如果该标签已经标记，则传入该参数；未标记，则不传。其中1为正常邮件，2为异常邮件。')
parser.add_argument('--category', type = str, default = '', help = '邮件类别(可选)，如果邮件类别已经标记，则传入该参数；未标记，则不传')
args = parser.parse_args()
path = args.path
valid_manual = args.manual
category = args.category

#建立数据库连接
db = pymysql.connect(host='114.55.139.105',user='root',password='Bit0123456789!',\
                     database='email-center',charset='utf8',port=3306)
#获取游标对象
cursor = db.cursor()

#插入数据语句
for file_path in traverse_path(path):
    char_set = "set names utf8mb4"
    cursor.execute(char_set)
    print(file_path)
    with open(file_path, 'rb') as fhdl:
        raw_email = fhdl.read()
    #编码替换，如果是GB2312编码，则换成GBK
#     print("+++++", raw_email)
#     print(chardet.detect(raw_email))
    raw_email_local = raw_email.decode('GBK', 'replace')
    if "GB2312" in raw_email_local:
        raw_email = raw_email_local.replace("GB2312", "GBK")
        raw_email = bytes(raw_email.encode('GBK', 'replace'))

    #解析邮件
    ep = eml_parser.EmlParser(include_raw_body = True)
    parsed_eml = ep.decode_email_bytes(raw_email)
    email_json = json.dumps(parsed_eml, default=json_serial, ensure_ascii=False)
#     print("-------", email_json)
    email_dict = json.loads(email_json)
    
#     print(parsed_eml)
    
    #email_from... 发件人
    email_from = email_dict['header']['from']
    #email_to...   收件人
    email_to = email_dict['header']['to']
    #email_date...  收件日期
    email_date = email_dict['header']['date']
    email_date = email_date.replace("T", " ")
    email_date = email_date.replace("+08:00", "")

    if "+00:00" in email_date:
        email_date = email_date.replace("+00:00", "")
        offset = datetime.timedelta(hours=8)
        email_date = datetime.datetime.strptime(email_date, '%Y-%m-%d %H:%M:%S')
        email_date = email_date + offset
    elif "+02:00" in email_date:
        email_date = email_date.replace("+02:00", "")
        offset = datetime.timedelta(hours=6)
        email_date = datetime.datetime.strptime(email_date, '%Y-%m-%d %H:%M:%S')
        email_date = email_date + offset
#     email_date = time.strftime("%Y-%m-%d %H:%M:%S", email_date)
    #email_subject... 邮件主题
    email_subject = email_dict['header']['subject']
    #email_content...   邮件正文
    for i in range(len(email_dict['body'])):
        if email_dict['body'][i]['content_type'] == "text/html":
            original_content = email_dict['body'][i]['content']
            break
        else:
            original_content = email_dict['body'][0]['content']
    original_content = str(original_content)
    #original_content_part局部变量
    original_content_part = original_content
    original_content_part = original_content_part.replace(" ", "")
    original_content_part = original_content_part.replace("\n", "")
    original_content_part = original_content_part.replace("\r", "")
    original_content_part = original_content_part.replace("\t", "")    
#     print(len(original_content))
    if original_content_part != "":
        email_content = html_parser(original_content)
    email_content = email_content.replace(" ", "")
    email_content = email_content.replace("\n", "")
    email_content = email_content.replace("\r", "")
    email_content = email_content.replace("\t", "")
    if len(email_content) > 10000:
        email_content = email_content[0:10000]
    #email_attachment...   邮件附件
    email_attachment = []
    if 'attachment' in email_dict:
        for i in email_dict['attachment']:
            email_attachment.append(i['filename'])
    file_path = file_path.replace(path + "\\", "")
    if '\\' in file_path:
        file_path = file_path.split('\\')[-1]
    else:
        file_path = file_path.split('/')[-1]
#     print(email_from)
#     print(email_to)
#     print(email_date)
#     print(email_subject)
#     print(email_attachment)
#     print(email_content)
    if valid_manual != 1 and valid_manual != 2 and category != '':
        query = """insert into body_test (file_name, body_test.from, send_time, \
        body_test.to, subject, category, content_length, attachments, body) values \
        ( '%s', '%s', '%s', '%s', '%s', '%s', %d, '%s', '%s')"""
        query = query%(pymysql.escape_string(str(file_path)), pymysql.escape_string(str(email_from)), email_date,\
                       pymysql.escape_string(str(email_to)), pymysql.escape_string(str(email_subject)), pymysql.escape_string(str(category)), int(len(email_content)), \
                       pymysql.escape_string(str(email_attachment)), pymysql.escape_string(str(email_content)))
    elif valid_manual != 1 and valid_manual != 2 and category == '':
        query = """insert into body_test (file_name, body_test.from, send_time, \
        body_test.to, subject, content_length, attachments, body) values \
        ( '%s', '%s', '%s', '%s', '%s', %d, '%s', '%s')"""
        query = query%(pymysql.escape_string(str(file_path)), pymysql.escape_string(str(email_from)), email_date,\
                       pymysql.escape_string(str(email_to)), pymysql.escape_string(str(email_subject)), int(len(email_content)), \
                       pymysql.escape_string(str(email_attachment)), pymysql.escape_string(str(email_content)))
    elif category == '' and valid_manual == 1 or valid_manual == 2:
        query = """insert into body_test (file_name, body_test.from, valid_manual, send_time, \
        body_test.to, subject, content_length, attachments, body) values \
        ( '%s', '%s', '%s', '%s', '%s', '%s', %d, '%s', '%s')"""
        query = query%(pymysql.escape_string(str(file_path)), pymysql.escape_string(str(email_from)), int(valid_manual), email_date,\
                       pymysql.escape_string(str(email_to)), pymysql.escape_string(str(email_subject)), int(len(email_content)), \
                       pymysql.escape_string(str(email_attachment)), pymysql.escape_string(str(email_content)))
    else:
        query = """insert into body_test (file_name, body_test.from, valid_manual, send_time, \
        body_test.to, subject, category, content_length, attachments, body) values \
        ( '%s', '%s', %d, '%s', '%s', '%s', '%s', %d, '%s', '%s')"""
        query = query%(pymysql.escape_string(str(file_path)), pymysql.escape_string(str(email_from)), int(valid_manual), email_date,\
                       pymysql.escape_string(str(email_to)), pymysql.escape_string(str(email_subject)), pymysql.escape_string(str(category)), int(len(email_content)), \
                       pymysql.escape_string(str(email_attachment)), pymysql.escape_string(str(email_content)))
#     print(query)
    cursor.execute(query)
    db.commit()
    
#关闭游标，提交，关闭数据库连接

#如果没有这些关闭操作，执行后在数据库中查看不到数据

cursor.close()

db.commit()

db.close()


# print(json.dumps(parsed_eml, default=json_serial))


#cursor.close()
#db.close()