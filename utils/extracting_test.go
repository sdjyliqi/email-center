package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//GetSenderDomain ...获取发件者的域名信息 customer_service@jd.com
func Test_GetFileNames(t *testing.T) {
	email, domain := "customer_service@jd.com", "jd"
	result := GetSenderDomain(email)
	t.Log(result, "======", domain)

	assert.Equal(t, domain, result)
	email, domain = "95555@message.cmbchina.com", "cmbchina"
	result = GetSenderDomain(email)
	t.Log(result, "======", domain)
	assert.Equal(t, domain, result)

	email, domain = "bitjsj@bit.edu.cn", "bit"
	result = GetSenderDomain(email)
	t.Log(result, "======", domain)
	assert.Equal(t, domain, result)
}

func Test_GetVX(t *testing.T) {
	content := "Vx:361212981你好1212，他的微信:sdjyliqi"
	result := GetVX(content)
	t.Log(result)
	//
	//content = "加v:361212981，"
	//result = GetVX(content)
	//t.Log(result)

	content = "1212asvx3612129as-81"
	result = GetVX(content)
	t.Log(result)
}

func Test_GetSenderDomain(t *testing.T) {
	content := "eos@t.cnbfund.com"
	expect := "cnbfund"
	assert.Equal(t, expect, GetSenderDomain(content))
}

func Test_GetQQ(t *testing.T) {
	content := "<html><body>+V：w13774336437   qq63123988</body></html>"
	result := GetQQ(content)
	assert.Equal(t, 1, len(result))
	t.Log(result)

	content = "加v:361212981，高清中文群p男女一夜请看懂加qq3023623334qcjbnhhyiorg"
	result = GetQQ(content)
	assert.Equal(t, 1, len(result))
	t.Log("++++++++++++", result)

	content = "<html><body>qq:361111888,+V：w13774336437   qq63123988</body></html>"
	result = GetQQ(content)
	assert.Equal(t, 2, len(result))
	t.Log(result)

	content = "可约，+扣3781281212"
	result = GetQQ(content)
	assert.Equal(t, 1, len(result))
	t.Log(result)

}

func Test_ExtractWebDomain(t *testing.T) {
	content := "美方在病毒溯源上“带节奏”极不负责,https://www.163.com/dy/article/GFS69GVG05346RC6.html,开具发票登录https://www.jd.com/?d"
	expect := []string{"163.com", "jd.com"}
	result, ok := ExtractWebDomain(content)
	t.Log(result, ok)
	assert.True(t, ok)
	assert.Equal(t, len(expect), len(result))
	assert.Equal(t, result[0], expect[0])
	assert.Equal(t, result[1], expect[1])

	content = "美方在病毒溯源上"
	result, ok = ExtractWebDomain(content)
	assert.False(t, ok)
}

func Test_ExtractMPhone(t *testing.T) {
	//for valid phone num
	phone := "正规13538046808薇同步发票"
	v, ok := ExtractMobilePhone(phone)
	t.Log(v, ok)
	//assert.True(t, ok)
	//assert.Equal(t, v, "15210510285")

	////for valid phone num
	//phone = "我的手机+8615210510285"
	//v, ok = ExtractMobilePhone(phone)
	//t.Log(v, ok)
	//assert.True(t, ok)
	//assert.Equal(t, v, "15210510285")
	//
	////for invalid phone num
	//phone = "152105"
	//v, ok = ExtractMobilePhone(phone)
	//t.Log(v, ok)
	//assert.False(t, ok)
}
