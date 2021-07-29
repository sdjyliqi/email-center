package utils

//定义邮件合法性标记
type LegalTag int8

const (
	UnknownTag LegalTag = 0 //无状态
	ValidTag   LegalTag = 1 //合法
	InvalidTag LegalTag = 2 //非法
)

var WebFormat = "(http|https)://[a-z0-9\\.]+"
var shortWebFormat = "[a-z0-9\\.]{2,12}.(cn|com)"
