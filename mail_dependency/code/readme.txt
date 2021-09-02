邮件解析脚本名称：email_parser_console.py
作用：遍历所选目录，解析邮件存入数据库
支持系统：支持Linux、Windows系统解析
环境依赖：python3、pip依赖包
示例：
python email_parser_console.py --path /usr/local/mail_dependency/email/
读取目录下所有文件，存入数据库中
参数：
  -h, --help           显示帮助信息
  --path PATH          邮件所在目录(必填)
  --manual MANUAL      是否为正常邮件(可选)，如果该标签已经标记，则传入该参数；未标记，则不传。其中1为正常邮件，2为异常邮件。
  --category CATEGORY  邮件类别(可选)，如果邮件类别已经标记，则传入该参数；未标记，则不传