package tryagain

import (
	"fmt"
	"regexp"
	"strconv"
)

//银行卡BIN前三位
var BandCardBIN = map[string]bool{
	"103":  true,
	"185":  true,
	"303":  true,
	"356":  true,
	"360":  true,
	"370":  true,
	"374":  true,
	"376":  true,
	"377":  true,
	"400":  true,
	"402":  true,
	"403":  true,
	"404":  true,
	"405":  true,
	"406":  true,
	"407":  true,
	"409":  true,
	"410":  true,
	"412":  true,
	"415":  true,
	"421":  true,
	"422":  true,
	"424":  true,
	"425":  true,
	"427":  true,
	"428":  true,
	"431":  true,
	"433":  true,
	"434":  true,
	"435":  true,
	"436":  true,
	"438":  true,
	"439":  true,
	"442":  true,
	"451":  true,
	"453":  true,
	"456":  true,
	"458":  true,
	"461":  true,
	"463":  true,
	"464":  true,
	"468":  true,
	"469":  true,
	"472":  true,
	"479":  true,
	"481":  true,
	"483":  true,
	"486":  true,
	"487":  true,
	"489":  true,
	"491":  true,
	"493":  true,
	"498":  true,
	"504":  true,
	"510":  true,
	"512":  true,
	"513":  true,
	"514":  true,
	"515":  true,
	"517":  true,
	"518":  true,
	"519":  true,
	"520":  true,
	"521":  true,
	"522":  true,
	"523":  true,
	"524":  true,
	"525":  true,
	"526":  true,
	"527":  true,
	"528":  true,
	"530":  true,
	"531":  true,
	"532":  true,
	"535":  true,
	"537":  true,
	"539":  true,
	"540":  true,
	"541":  true,
	"543":  true,
	"544":  true,
	"545":  true,
	"547":  true,
	"548":  true,
	"549":  true,
	"550":  true,
	"552":  true,
	"553":  true,
	"554":  true,
	"556":  true,
	"557":  true,
	"558":  true,
	"559":  true,
	"566":  true,
	"584":  true,
	"589":  true,
	"601":  true,
	"602":  true,
	"603":  true,
	"609":  true,
	"620":  true,
	"621":  true,
	"622":  true,
	"623":  true,
	"624":  true,
	"625":  true,
	"626":  true,
	"627":  true,
	"628":  true,
	"629":  true,
	"644":  true,
	"650":  true,
	"682":  true,
	"683":  true,
	"685":  true,
	"688":  true,
	"690":  true,
	"694":  true,
	"695":  true,
	"843":  true,
	"870":  true,
	"888":  true,
	"900":  true,
	"905":  true,
	"909":  true,
	"911":  true,
	"920":  true,
	"921":  true,
	"940":  true,
	"941":  true,
	"947":  true,
	"955 ": true,
	"966":  true,
	"968":  true,
	"984 ": true,
	"985":  true,
	"989":  true,
	"990":  true,
	"998":  true,
	"999":  true,
}

//PickBankCard ...获取银行卡号
func PickBankCard(txt string) []string {
	var bankCard []string
	//银行卡首位没有0、2、7
	//18位银行卡和身份证号重复，暂不匹配
	pattern := `[^\d.][1345689]([0-9]{13,16}|[0-9]{18})[^\d.]`
	reg := regexp.MustCompile(pattern)
	if reg == nil {
		fmt.Println("regexp err")
	}
	cardIDList := reg.FindAllString(txt, -1)
	//判断正则匹配的银行卡号切片是否为空，如果为空则空切片
	if cardIDList == nil {
		return bankCard
	}
	//对提取出的银行卡号用Luhn算法校验
	for _, value := range cardIDList {
		value = value[1 : len(value)-1]
		ok := BandCardBIN[value[:3]]
		if !ok {
			continue
		}
		sum := 0
		//第一步，从卡号最后一位数字开始，逆向将奇数位(1、3、5等等)相加
		for i := len(value) - 1; i >= 0; i -= 2 {
			num, _ := strconv.Atoi(string(value[i]))
			sum += num
		}
		//第二步，从卡号最后一位数字开始，逆向将偶数位数字，先乘以2（如果乘积为两位数，则将其减去9），再求和。
		for i := len(value) - 2; i >= 0; i -= 2 {
			num, _ := strconv.Atoi(string(value[i]))
			if (num * 2) > 9 {
				sum += num*2 - 9
			} else {
				sum += num * 2
			}
		}
		//第三步，将奇数位总和加上偶数位总和，结果应该可以被10整除
		if sum%10 == 0 {
			bankCard = append(bankCard, value)
		}
	}
	return bankCard
}
