package tryagain

import (
	"regexp"
	"strconv"
)

var provinceInfo = map[string]string{
	"11":  "北京市",
	"12":  "天津市",
	"13":  "河北省",
	"14":  "山西省",
	"15":  "内蒙古自治区",
	"21":  "辽宁省",
	"22":  "吉林省",
	"23":  "黑龙江省",
	"31":  "上海市",
	"32":  "江苏省",
	"33":  "浙江省",
	"34":  "安徽省",
	"35":  "福建省",
	"36":  "江西省",
	"37":  "山东省",
	"41":  "河南省",
	"42":  "湖北省",
	"43":  "湖南省",
	"44":  "广东省",
	"45":  "广西壮族自治区",
	"46":  "海南省",
	"50":  "重庆市",
	"51":  "四川省",
	"52":  " 贵州省",
	"53":  "云南省",
	"54":  "西藏自治区",
	"61 ": "陕西省",
	"62":  "甘肃省",
	"63":  "青海省",
	"64":  "宁夏回族自治区",
	"65":  "新疆维吾尔自治区",
	"71":  "台湾省",
	"81":  "香港特别行政区",
	"82":  "澳门特别行政区",
}

var LastDigital = map[int]string{
	0:  "1",
	1:  "0",
	2:  "X",
	3:  "9",
	4:  "8",
	5:  "7",
	6:  "6",
	7:  "5",
	8:  "4",
	9:  "3",
	10: "2",
}

/*
* 13开头排序：(0-9)（134 135 136 137 138 139 130 131 132 133）
* 14开头排序：(5-9)（147 148 145 146 149）
* 15开头排序：(0-3|5-9)（150 151 152 157 158 159 155 156 153）
* 16开头排序：(6-7)（166 167）
* 17开头排序：(1-8)（172 178 171 175 176 173 174 177）
* 18开头排序：(0-9)（182 183 184 187 188 185 186 180 181 189）
* 19开头排序：(1|3|5|6|8|9)（195 198 196 191 193 199）
 */
var MobilePhonePrefix = map[string]bool{
	"130": true,
	"131": true,
	"132": true,
	"133": true,
	"134": true,
	"135": true,
	"136": true,
	"137": true,
	"138": true,
	"139": true,
	"145": true,
	"146": true,
	"147": true,
	"148": true,
	"149": true,
	"150": true,
	"151": true,
	"152": true,
	"153": true,
	"155": true,
	"156": true,
	"157": true,
	"158": true,
	"159": true,
	"166": true,
	"167": true,
	"171": true,
	"172": true,
	"173": true,
	"174": true,
	"175": true,
	"176": true,
	"177": true,
	"178": true,
	"181": true,
	"182": true,
	"183": true,
	"184": true,
	"185": true,
	"186": true,
	"187": true,
	"188": true,
	"189": true,
	"191": true,
	"193": true,
	"195": true,
	"196": true,
	"198": true,
	"199": true,
}

func ExtractCardIDs(txt string) []string {
	var cardIDs []string
	pattern := `[1-8]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]`
	reg := regexp.MustCompile(pattern)
	cardIDList := reg.FindAllString(txt, -1)
	weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	for _, v := range cardIDList {
		cardScore := 0
		if len(v) != 18 {
			continue
		}
		_, ok := provinceInfo[v[0:2]]
		if !ok {
			continue
		}
		for k, vv := range []rune(v)[0:17] {
			digitalNum, _ := strconv.Atoi(string(vv))
			cardScore += weight[k] * digitalNum
		}
		sign, _ := LastDigital[cardScore%11]
		if sign == v[17:] {
			cardIDs = append(cardIDs, v)
		}
	}
	return cardIDs
}

/*
 * 13开头排序：(0-9)（134 135 136 137 138 139 130 131 132 133）
 * 14开头排序：(5-9)（147 148 145 146 149）
 * 15开头排序：(0-3|5-9)（150 151 152 157 158 159 155 156 153）
 * 16开头排序：(6-7)（166 167）
 * 17开头排序：(1-8)（172 178 171 175 176 173 174 177）
 * 18开头排序：(0-9)（182 183 184 187 188 185 186 180 181 189）
 * 19开头排序：(1|3|5|6|8|9)（195 198 196 191 193 199）
 */

func ExtractMobilePhoneDs(txt string) []string {
	var phoneIDs []string
	var PhoneFormat = "[^0-9.-_](1[3-9]\\d{9})[^0-9]"
	phoneRegx := regexp.MustCompile(PhoneFormat)
	phoneNums := phoneRegx.FindAllString(txt, -1)
	for _, v := range phoneNums {
		phoneInfo := []rune(v)
		vLen := len(phoneInfo)
		startIndex, stopIdx := 0, len(phoneInfo)
		if phoneInfo[0] < '0' || phoneInfo[0] > '9' {
			startIndex = 1
		}
		if phoneInfo[vLen-1] < '0' || phoneInfo[vLen-1] > '9' {
			stopIdx = stopIdx - 1
		}
		phoneID := string(phoneInfo[startIndex:stopIdx])
		phoneIDPrefix := phoneID[0:3]
		_, ok := MobilePhonePrefix[phoneIDPrefix]
		if ok {
			phoneIDs = append(phoneIDs, phoneID)
		}
	}
	return phoneIDs
}

//从里面抽离地址信息
func ExtractAddr(txt string) []string {
	//

	return nil
}
