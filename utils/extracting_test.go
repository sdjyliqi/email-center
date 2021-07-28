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
