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
	content := "vx:361212981你好1212，他的微信:sdjyliqi"
	result := GetVX(content)
	t.Log(result)

	content = "加v:361212981，"
	result = GetVX(content)
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
