package tryagain

import (
	"email-center/ac"
	"regexp"
	"strings"
)

//PickAddr ...获取地址信息
func PickAddr(content string) []string {
	var addrList []string
	addrFormat := `(北京|天津|上海|重庆|河北|河南|云南|辽宁|黑龙江|湖南|安徽|山东|新疆维吾尔| 新疆|江苏|浙江|江西|湖北|广西壮族|广西|甘肃|山西|内蒙古| 陕西|吉林|福建|贵州|广东|青海|西藏|四川省|宁夏回族|海南| )+(市|省|自治区)+[^ ，,。.!]+`
	addrRegx := regexp.MustCompile(addrFormat)
	addrRaws := addrRegx.FindAllString(content, -1)
	for _, v := range addrRaws {
		result := ac.GetAddrWords(v)
		if len(result) > 0 && strings.HasPrefix(v, result[0]) && len(v) > len(result[0])+5 {
			addrList = append(addrList, v)
		}
	}
	return addrList
}
