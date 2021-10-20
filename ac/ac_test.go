package ac

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

	result, _ = GetCategoryIdx("小哥哥，老公不在家，可约，可全套，校花，萝莉、空姐")
	t.Log(result)

}

func Test_IGetWhiteHighlights(t *testing.T) {
	result := GetWhiteHighlights("登录jd.com,官方客服电话950618")
	t.Log(result)
}
