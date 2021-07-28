package utils

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
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

func PickupEmail(path string, validTag LegalTag) (*Email, error) {
	fmt.Println("----------PickupEmail start------------------")
	fromPrefix, toPrefix, msgIDPrefix, subjectPrefix, encodePrefix := "From:", "To:", "Message-ID:<", "Subject:", "Content-Transfer-Encoding:"
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
		line := strings.Replace(v, " ", "", -1)
		line = strings.ReplaceAll(line, "\n", "")
		switch {
		case strings.HasPrefix(line, fromPrefix):
			//优先提取<>中内内容，如果没有直接获取全部内容
			start := strings.Index(line, "<")
			stop := strings.Index(line, ">")
			if start >= 0 && start < stop {
				from := line[start+1 : stop]
				pickupInfo.From = from
			} else {
				from := line[len(toPrefix)+1:]
				pickupInfo.From = from
			}
		case strings.HasPrefix(line, toPrefix):
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
			sendTime := line[len("Date:"):]
			pickupInfo.Date = sendTime
		case strings.HasPrefix(line, subjectPrefix):
			subject := line[len(subjectPrefix) : len(line)-1]
			pickupInfo.Subject = subject

		case strings.HasPrefix(line, msgIDPrefix):
			msgID := line[len(msgIDPrefix) : len(line)-1]
			pickupInfo.MessageID = msgID
		//处理邮件正文编码
		case strings.HasPrefix(line, encodePrefix):
			encoding := line[len(encodePrefix):]
			fmt.Println("------------encoding-----------------------", encoding)
			pickupInfo.Encoding = encoding
		}
	}
	if pickupInfo.Encoding == "base64" {
		decodeContent, err := DecodingBase64(pickupInfo.ContentBody)
		if err == nil {
			pickupInfo.ContentBody = decodeContent
		}
	}
	aaa, _ := json.Marshal(pickupInfo)
	fmt.Println("----------PickupEmail end------------------", string(aaa))
	return pickupInfo, err
}
