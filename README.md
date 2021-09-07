邮件解析可执行包使用说明
版本：  email_parser_windows.exe为windows版本
        email_parser_linux为linux版本
作用：遍历本地目录中的邮件，解析完成后写入数据库中
示例：
Windows：
email_parser_windows.exe -a C:\Users\user\Desktop\创新项目\垃圾邮件\test -a C:\Users\user\Desktop\创新项目\垃圾邮件\test -t body_test
 
Linux：
./email_parser_linux -a test/ -t body_test
注：先给email_parser_linux赋予执行权限 chmod 766 email_parser_linux
 
参数： 
  -h, --help            show this help message and exit
  -a PATH, --path PATH  邮件所在目录(必填)
  -m MANUAL, --manual MANUAL
     是否为正常邮件(可选)，如果该标签已经标记，则传入该参数；未标记，则不传。其中1为正常邮件，2为异常邮件。
  -c CATEGORY, --category CATEGORY
    邮件类别(可选，默认为空)，如果邮件类别已经标记，则传入该参数；未标记，则不传
  -t TABLE, --table TABLE
    表名称（可选，默认为body）
  -H HOST, --host HOST  数据库IP（可选，默认为10.233.146.47）
  -P PORT, --port PORT  数据库端口（可选，默认为16315）
  -d DATABASE, --database DATABASE
      数据库名称（可选，默认为email-center）
  -u USER, --user USER  数据库用户名（可选，默认为root）
  -p PASSWORD, --password PASSWORD
       数据库密码（可选）
