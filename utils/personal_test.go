package utils

import (
	"regexp"
	"testing"
	"time"
)

//身份证号的长度是18位置
func Test_CardID(t *testing.T) {

	t.Log(time.Now())
	text := "我们在使用分列功能进行出生日期提取之前，0.152102198510281834先来看看身份证号码的组成：二代身份证都是18位，前面6位152102198510281834中文是地区代码，中间8位是出生年月日，后面三位是编码，最后一位是识别码。身份证号012152102198510281834中文"
	ids := ExtractCardIDs(text)
	t.Log(ids)
	t.Log(time.Now())
}

func Test_PhoneID(t *testing.T) {
	text := "我0.15210510987们在使14012346789用分列功能进1521051028578行出生日期提取之前，我手机号15210510285,0.115210510285，   18701516837开心"
	v := ExtractMobilePhoneDs(text)
	t.Log(v)
}

func Test_Addr(t *testing.T) {
	text := `籍贯是山东省巨野县。我的单位地址是北京市海淀区汇通路12号中国光大银行1层，
         现居住地址为北京市丰台区太平桥西里11#907， 
        老家地址为河北省石家庄市灵寿县塔上镇万里村，那边比较穷" +
		"我买的房子地址是北京市丰台区郑尚名苑1号楼1单位804`
	//addrFormat:= `([^省]+省|.+自治区|[^市]+市)([^自治州]+自治州|[^市]+市|[^盟]+盟|[^地区]+地区|.+区划)([^市]+市|[^县]+县|[^旗]+旗|.+区)`
	addrFormat := `(北京|天津|河北|山东)+(市|省|自治区)+[^ ，,。.!]+`
	phoneRegx := regexp.MustCompile(addrFormat)
	phoneNums := phoneRegx.FindAllString(text, -1)
	t.Log(phoneNums)
	//然后里面包括省-市-县  和 省--县构建的ac自动机，并且不仅包括这些信息。

}
