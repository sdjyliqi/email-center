package utils

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

//ExtractMPhone ..提取手机号

//GetSenderDomain ...获取发件者的域名信息 customer_service@jd.com
//处理步骤大致如下，先抽取@符号后面的内容，然后在获取com之前的，然后用.进行分割，取后面的
//bitjsj@bit.edu.cn
func GetSenderDomain(email string) string {
	content := strings.Trim(email, " ")
	idx := strings.Index(email, "@")
	if idx > 1 {
		content = content[idx+1:]
	}
	//处理一般性商业邮箱， .com.cn在这儿处理掉
	idx = strings.Index(content, ".com")
	if idx > 1 {
		content = content[:idx]
	}
	//处理教育邮箱
	idx = strings.Index(content, ".edu.")
	if idx > 1 {
		content = content[:idx]
	}
	//处理无.com，只有.cn的情况   styy@tpp.cntaiping.com对于这种情况使用条件 idx+3==len(content) 来排除
	idx = strings.Index(content, ".cn")
	if idx > 1 && idx+3 == len(content) {
		content = content[:idx]
	}
	if content == "" {
		return ""
	}
	info := strings.Split(content, ".")
	return info[len(info)-1]
}

//GetQQ ...获取QQ号
func GetQQ(content string) []string {
	var QQIDs []string
	content = content + " "
	content = strings.Replace(content, ":", "", -1)
	content = strings.Replace(content, "：", "", -1)
	var qqFormat = "(q|扣|抠|旧)[0-9]{5,11}[^@A-Za-z]"
	formatRegx := regexp.MustCompile(qqFormat)
	values := formatRegx.FindAllStringSubmatch(strings.ToLower(content), -1)
	for _, v := range values {
		QQIDs = append(QQIDs, v[0])
	}
	return QQIDs
}

//GetQQ ...获取QQ号
func GetVX(content string) []string {
	var weixinIDs []string
	var qqFormat = "(加v:|加v|\\+v：|\\+v:|vx|vx:|vx：|微信|微信:|微信：)[a-z0-9-_]{6,64}"
	formatRegx := regexp.MustCompile(qqFormat)
	values := formatRegx.FindAllStringSubmatch(strings.ToLower(content), -1)
	for _, v := range values {
		idx := strings.Index(content, v[0])
		fmt.Println(idx)
		if idx > 0 {
			formerLetter := rune(content[idx-1])
			if unicode.IsLetter(formerLetter) {
				continue
			}
		}
		weixinIDs = append(weixinIDs, v[0])
	}
	return weixinIDs
}

//ExtractWebDomain ..提取登录网址
func ExtractWebDomain(txt string) ([]string, bool) {
	txt = strings.ToLower(txt)
	var domains []string
	formatRegx := regexp.MustCompile(WebFormat)
	values := formatRegx.FindAllStringSubmatch(txt, -1)
	if len(values) == 0 {
		return nil, false
	}
	for _, v := range values {
		//todo 后续替换这个再进行优化吧
		url := strings.Replace(v[0], "https://www.", "", 1)
		url = strings.Replace(url, "http://www.", "", 1)
		url = strings.Replace(url, "https://", "", 1)
		url = strings.Replace(url, "http://", "", 1)
		domains = append(domains, url)
	}
	return domains, true
}

//ExtractMPhone ..提取手机号
func ExtractMobilePhone(txt string) ([]string, bool) {
	txt = txt + " "
	var ids []string
	txt = strings.Replace(txt, "-", "", -1)
	txt = strings.Replace(txt, "+86", "", -1)
	phoneRegx := regexp.MustCompile(PhoneFormat)
	phoneNums := phoneRegx.FindAllStringSubmatch(txt, -1)
	for _, v := range phoneNums {
		if len(v[0]) > lenMobilePhone {
			ids = append(ids, v[0][:lenMobilePhone])
		}
	}
	return ids, len(ids) > 0
}

func ChkContentIsMobilePhone(txt string) bool {
	txt = strings.Replace(txt, "+86", "", 1)
	if len(txt) != lenMobilePhone {
		return false
	}
	_, ok := ExtractMobilePhone(txt)
	return ok
}
