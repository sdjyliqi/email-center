package utils

import (
	"testing"
)

///Test_Base64 ... 测试base64 加解密的算法
func Test_InitURLDomainAC(t *testing.T) {
	var tags []string
	InitURLDomainAC()
	t.Log(DomainACMatch)
	result := DomainACMatch.Match("http://jd.com,欢迎登录")
	t.Log(result)
	for _, v := range result {
		tags = append(tags, URLDomains[v])
	}
	t.Log(tags)
}
func Test_GetCategoryIdx(t *testing.T) {
	result, _ := GetCategoryIdx("发票，电子发票。kai发票")
	t.Log(result)
}
