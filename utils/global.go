package utils

//定义邮件合法性标记
type LegalTag int8

type Category int8

const (
	UnknownTag     LegalTag = 0  //无状态
	ValidTag       LegalTag = 1  //合法
	InvalidTag     LegalTag = 2  //非法
	lenMobilePhone int      = 11 //移动电话标准长度

	UnknownCategory Category = 0
	BillCategory    Category = 1 //发票类
	AdvertCategory  Category = 2 //广告类
)

func (c Category) Name() string {
	switch c {
	case BillCategory:
		return "发票类"
	case AdvertCategory:
		return "广告类"
	default:
		return "未知分类"
	}
}

var WebFormat = "(http|https)://[a-z0-9\\.]+"
var shortWebFormat = "[a-z0-9\\.]{2,12}.(cn|com)"

var PhoneFormat = "(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}" //手机号码格式
var TimeFormat = "2006-01-02 15:04:05"

type Charset string

var TagProperty = map[string]LegalTag{
	"发票": UnknownTag,
	"發票": InvalidTag,
	"潑镖": InvalidTag,
}
