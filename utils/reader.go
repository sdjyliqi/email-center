package utils

import (
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"regexp"
	"strings"
)

type Email struct {
	From            string   //发件人
	To              string   //收件人
	Subject         string   //邮件主题
	Date            string   //发生时间
	MessageID       string   //消息id
	ContentLanguage string   //消息语言类型
	ContentBody     string   //邮件正文
	Category        string   //邮件分类
	FileName        string   //邮件文件路径
	Encoding        string   //邮件正文编码
	Valid           LegalTag //邮件合法标记
}

func GetFileNames(path, dot string) ([]string, error) {
	var targetFiles []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("read dir error")
		return targetFiles, err
	}
	for i, v := range files {
		fmt.Println(i, "=", v.Name())
		fullName := fmt.Sprintf("%s/%s", path, v.Name())
		if v.Name() == "" {
			continue
		}
		if dot != "" {
			fileInfo := strings.Split(v.Name(), ".")
			if len(fileInfo) == 2 && fileInfo[1] == dot {
				targetFiles = append(targetFiles, fullName)
			}
		} else {
			targetFiles = append(targetFiles, fullName)
		}
	}
	return targetFiles, nil
}

func ReadEmail(path string) ([]byte, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		glog.Errorf("Open file %s by ioutil.ReadFile failed,err:%+v", path, err)
		return nil, err
	}
	return contents, err
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func pickupFromEmail(line string) string {
	idx, from := "From:", ""
	start := strings.Index(line, "<")
	stop := strings.Index(line, ">")
	if start >= 0 && start < stop {
		from = line[start+1 : stop]
	} else {
		from = line[len(idx)+1:]
	}
	ok := VerifyEmailFormat(from)
	if ok {
		return from
	}
	return ""
}

//
func PickupEmailBody(content string) string {
	body, keyIndex, delIndex := "", "Content-Type: text/plain; charset=utf-8", "Content-Transfer-Encoding: base64"
	idxStart := strings.Index(content, keyIndex)
	if idxStart > 0 {
		content = content[idxStart+len(keyIndex):]
	} else {
		//寻找第二标记 Content-Type: text/html; charset=utf-8
		idxStart = strings.Index(content, "Content-Type: text/html; charset=utf-8")
		if idxStart > 0 {
			content = content[idxStart+len(keyIndex):]
		}
	}
	//寻找结束的位置
	idxStop := strings.Index(content, "----boundary_")
	body = content
	if idxStop > 0 {
		body = content[0:idxStop]
	}
	body = strings.ReplaceAll(body, delIndex, "")
	body = strings.ReplaceAll(body, "\r\n", "")
	body, _ = DecodingBase64(body)
	fmt.Println(body)
	return body
}

func PickupEmail(path string, validTag LegalTag) (*Email, error) {
	fromPrefix, toPrefix, msgIDPrefix, subjectPrefix, encodePrefix := "From:", "To:", "Message-ID:<", "Subject:", "Content-Transfer-Encoding:"
	pickupInfo := &Email{}
	content, err := ReadEmail(path)
	if err != nil {
		return nil, err
	}
	idx := strings.Index(string(content), "\r\n\r")
	header := string(content)[0:idx]
	pickupInfo.ContentBody = PickupEmailBody(string(content))

	info := strings.Split(header, "\r")
	partIDX := ""
	for _, v := range info {
		line := strings.Replace(v, " ", "", -1)
		line = strings.ReplaceAll(line, "\n", "")
		switch {
		case strings.HasPrefix(line, fromPrefix):
			partIDX = fromPrefix
			//优先提取<>中内内容，如果没有直接获取全部内容
			pickupInfo.From = pickupFromEmail(line)
		case strings.HasPrefix(line, toPrefix):
			partIDX = toPrefix
			//提取<>中内内容,如果没有直接获取全部内容
			start := strings.Index(line, "<")
			stop := strings.Index(line, ">")
			if start >= 0 && start < stop {
				to := line[start+1 : stop]
				pickupInfo.To = to
			} else {
				to := line[len(toPrefix)+1:]
				pickupInfo.To = to
			}
		case strings.HasPrefix(line, "Date:"):
			partIDX = "Date:"
			sendTime := line[len("Date:"):]
			pickupInfo.Date = sendTime
		case strings.HasPrefix(line, subjectPrefix):
			partIDX = subjectPrefix
			subject := line[len(subjectPrefix) : len(line)-1]
			pickupInfo.Subject = subject

		case strings.HasPrefix(line, msgIDPrefix):
			partIDX = line
			msgID := line[len(msgIDPrefix) : len(line)-1]
			pickupInfo.MessageID = msgID
		//处理邮件正文编码
		case strings.HasPrefix(line, encodePrefix):
			partIDX = line
			encoding := line[len(encodePrefix):]
			pickupInfo.Encoding = encoding
		default:
			if partIDX == fromPrefix {
				line = fromPrefix + line
				senderID := pickupFromEmail(line)
				if senderID != "" {
					pickupInfo.From = senderID
				}
			}
		}
	}
	return pickupInfo, err
}
