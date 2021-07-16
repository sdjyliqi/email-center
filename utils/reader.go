package utils

import (
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"strings"
)

type Email struct {
	From            string //发件人
	To              string //收件人
	Subject         string //邮件主题
	Date            string //发生时间
	MessageID       string //消息id
	ContentLanguage string //消息语言类型
	ContentBody     string //邮件正文
	Category        string //邮件分类
	FileName        string //邮件文件路径
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

func PickupEmail(path string) (*Email, error) {
	fromPrefix, toPrefix, msgIDPrefix, subjectPrefix := "From: <", "To: <", "Message-ID: <", "Subject:"
	pickupInfo := &Email{}
	content, err := ReadEmail(path)
	if err != nil {
		return nil, err
	}
	idx := strings.Index(string(content), "\r\n\r")

	header := string(content)[0:idx]

	body := string(content)[idx:len(string(content))]
	pickupInfo.ContentBody = body

	info := strings.Split(header, "\r")

	for _, v := range info {
		line := strings.Trim(v, " ")
		line = strings.ReplaceAll(line, "\n", "")
		//t.Log(k,line)
		switch {
		case strings.HasPrefix(line, fromPrefix):
			from := line[len(fromPrefix) : len(line)-1]
			pickupInfo.From = from
		case strings.HasPrefix(line, toPrefix):
			to := line[len(toPrefix) : len(line)-1]
			pickupInfo.To = to
		case strings.HasPrefix(line, "Date:"):
			sendTime := line[:len(line)-1]
			pickupInfo.Date = sendTime
		case strings.HasPrefix(line, subjectPrefix):
			subject := line[len(subjectPrefix) : len(line)-1]
			pickupInfo.Subject = subject

		case strings.HasPrefix(line, msgIDPrefix):
			msgID := line[len(msgIDPrefix) : len(line)-1]
			pickupInfo.MessageID = msgID

		}
	}
	return pickupInfo, err
}
