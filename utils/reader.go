package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DusanKasan/parsemail"
	"github.com/golang/glog"
	"io/ioutil"
	"regexp"
	"strings"
)

type Email struct {
	Category   string   //邮件分类
	FileName   string   //邮件文件路径
	Encoding   string   //邮件正文编码
	Valid      LegalTag //邮件合法标记
	Body       string
	ParseEmail *parsemail.Email
}

func GetFileNames(path, dot string) ([]string, error) {
	var targetFiles []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("read dir error")
		return targetFiles, err
	}
	for _, v := range files {
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

func ParseEmailByPath(path string) (*parsemail.Email, error) {
	contents, err := ioutil.ReadFile(path)
	//reader.Read(contents)
	emailReader := bytes.NewReader(contents)
	fmt.Println(emailReader)
	email, err := parsemail.Parse(emailReader) // returns Email struct and error
	if err != nil {
		glog.Errorf("parsemail.Parse failed,err:%+v", err)
		return nil, err
	}
	return &email, nil
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

func PickupEmail(path string, validTag LegalTag) (*Email, error) {
	pickupInfo := &Email{Valid: validTag}
	parseValue, err := ParseEmailByPath(path)
	if err != nil {
		return nil, err
	}
	pickupInfo.ParseEmail = parseValue
	ccc, _ := json.Marshal(parseValue)
	fmt.Println(parseValue.Date, string(ccc))
	encoding, ok := pickupInfo.ParseEmail.Header["Content-Transfer-Encoding"]
	if ok && len(encoding) > 0 && encoding[0] == "base64" {
		encodingContent := pickupInfo.ParseEmail.TextBody
		if encodingContent == "" {
			encodingContent = pickupInfo.ParseEmail.HTMLBody
		}
		decodeBody, err := DecodingBase64(encodingContent)
		if err != nil {
			return nil, err
		}
		pickupInfo.Body = decodeBody
	}
	sendTime, _ := parseValue.Header.Date()
	fmt.Println("===========", sendTime.Format(TimeFormat))
	return pickupInfo, err
}
