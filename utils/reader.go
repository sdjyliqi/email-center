package utils

import (
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"regexp"
	"strings"
)

type Email struct {
	Category    string   //邮件分类
	FileName    string   //邮件文件路径
	Encoding    string   //邮件正文编码
	Valid       LegalTag //邮件合法标记
	Body        string
	Attachments []string //附件的名称
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
