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

var PhoneFormat = "(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}[^@A-Za-z]" //手机号码格式
var TimeFormat = "2006-01-02 15:04:05"
var SMTP163Host = "smtp.163.com:25"

//定义发票分类的关键字，通过关键字可能会判断出是否为异常短信
var TagProperty = map[string]LegalTag{
	"开具发票":      UnknownTag,
	"kaifapiao": InvalidTag,
	"kaifa票":    InvalidTag,
	"kaifa瞟":    InvalidTag,
	"kaifa漂":    InvalidTag,
	"kaifa镖":    InvalidTag,
	"kaifa標":    InvalidTag,
	"kaifa磦":    InvalidTag,
	"kaifa骠":    InvalidTag,
	"kaifa嘌":    InvalidTag,
	"kaifa缥":    InvalidTag,
	"kaifa螵":    InvalidTag,
	"kaifa彯":    InvalidTag,
	"kaifa慓":    InvalidTag,
	"kaifa剽":    InvalidTag,
	"kaifa徱":    InvalidTag,
	"kaifa僄":    InvalidTag,
	"kaifa飘":    InvalidTag,
	"kai发piao":  InvalidTag,
	"kai发票":     InvalidTag,
	"kai发瞟":     InvalidTag,
	"kai发漂":     InvalidTag,
	"kai发镖":     InvalidTag,
	"kai发標":     InvalidTag,
	"kai发磦":     InvalidTag,
	"kai发骠":     InvalidTag,
	"kai发嘌":     InvalidTag,
	"kai发缥":     InvalidTag,
	"kai发螵":     InvalidTag,
	"kai发彯":     InvalidTag,
	"kai发慓":     InvalidTag,
	"kai发剽":     InvalidTag,
	"kai发徱":     InvalidTag,
	"kai发僄":     InvalidTag,
	"kai发飘":     InvalidTag,
	"kai發piao":  InvalidTag,
	"kai發票":     InvalidTag,
	"kai發瞟":     InvalidTag,
	"kai發漂":     InvalidTag,
	"kai發镖":     InvalidTag,
	"kai發標":     InvalidTag,
	"kai發磦":     InvalidTag,
	"kai發骠":     InvalidTag,
	"kai發嘌":     InvalidTag,
	"kai發缥":     InvalidTag,
	"kai發螵":     InvalidTag,
	"kai發彯":     InvalidTag,
	"kai發慓":     InvalidTag,
	"kai發剽":     InvalidTag,
	"kai發徱":     InvalidTag,
	"kai發僄":     InvalidTag,
	"kai發飘":     InvalidTag,
	"kai潑piao":  InvalidTag,
	"kai潑票":     InvalidTag,
	"kai潑瞟":     InvalidTag,
	"kai潑漂":     InvalidTag,
	"kai潑镖":     InvalidTag,
	"kai潑標":     InvalidTag,
	"kai潑磦":     InvalidTag,
	"kai潑骠":     InvalidTag,
	"kai潑嘌":     InvalidTag,
	"kai潑缥":     InvalidTag,
	"kai潑螵":     InvalidTag,
	"kai潑彯":     InvalidTag,
	"kai潑慓":     InvalidTag,
	"kai潑剽":     InvalidTag,
	"kai潑徱":     InvalidTag,
	"kai潑僄":     InvalidTag,
	"kai潑飘":     InvalidTag,
	"kai泼piao":  InvalidTag,
	"kai泼票":     InvalidTag,
	"kai泼瞟":     InvalidTag,
	"kai泼漂":     InvalidTag,
	"kai泼镖":     InvalidTag,
	"kai泼標":     InvalidTag,
	"kai泼磦":     InvalidTag,
	"kai泼骠":     InvalidTag,
	"kai泼嘌":     InvalidTag,
	"kai泼缥":     InvalidTag,
	"kai泼螵":     InvalidTag,
	"kai泼彯":     InvalidTag,
	"kai泼慓":     InvalidTag,
	"kai泼剽":     InvalidTag,
	"kai泼徱":     InvalidTag,
	"kai泼僄":     InvalidTag,
	"kai泼飘":     InvalidTag,
	"kai蕟piao":  InvalidTag,
	"kai蕟票":     InvalidTag,
	"kai蕟瞟":     InvalidTag,
	"kai蕟漂":     InvalidTag,
	"kai蕟镖":     InvalidTag,
	"kai蕟標":     InvalidTag,
	"kai蕟磦":     InvalidTag,
	"kai蕟骠":     InvalidTag,
	"kai蕟嘌":     InvalidTag,
	"kai蕟缥":     InvalidTag,
	"kai蕟螵":     InvalidTag,
	"kai蕟彯":     InvalidTag,
	"kai蕟慓":     InvalidTag,
	"kai蕟剽":     InvalidTag,
	"kai蕟徱":     InvalidTag,
	"kai蕟僄":     InvalidTag,
	"kai蕟飘":     InvalidTag,
	"kai犮piao":  InvalidTag,
	"kai犮票":     InvalidTag,
	"kai犮瞟":     InvalidTag,
	"kai犮漂":     InvalidTag,
	"kai犮镖":     InvalidTag,
	"kai犮標":     InvalidTag,
	"kai犮磦":     InvalidTag,
	"kai犮骠":     InvalidTag,
	"kai犮嘌":     InvalidTag,
	"kai犮缥":     InvalidTag,
	"kai犮螵":     InvalidTag,
	"kai犮彯":     InvalidTag,
	"kai犮慓":     InvalidTag,
	"kai犮剽":     InvalidTag,
	"kai犮徱":     InvalidTag,
	"kai犮僄":     InvalidTag,
	"kai犮飘":     InvalidTag,
	"开fapiao":   InvalidTag,
	"开fa票":      InvalidTag,
	"开fa瞟":      InvalidTag,
	"开fa漂":      InvalidTag,
	"开fa镖":      InvalidTag,
	"开fa標":      InvalidTag,
	"开fa磦":      InvalidTag,
	"开fa骠":      InvalidTag,
	"开fa嘌":      InvalidTag,
	"开fa缥":      InvalidTag,
	"开fa螵":      InvalidTag,
	"开fa彯":      InvalidTag,
	"开fa慓":      InvalidTag,
	"开fa剽":      InvalidTag,
	"开fa徱":      InvalidTag,
	"开fa僄":      InvalidTag,
	"开fa飘":      InvalidTag,
	"开发piao":    InvalidTag,
	"开发票":       UnknownTag,
	"开发瞟":       InvalidTag,
	"开发漂":       InvalidTag,
	"开发镖":       InvalidTag,
	"开发標":       InvalidTag,
	"开发磦":       InvalidTag,
	"开发骠":       InvalidTag,
	"开发嘌":       InvalidTag,
	"开发缥":       InvalidTag,
	"开发螵":       InvalidTag,
	"开发彯":       InvalidTag,
	"开发慓":       InvalidTag,
	"开发剽":       InvalidTag,
	"开发徱":       InvalidTag,
	"开发僄":       InvalidTag,
	"开发飘":       InvalidTag,
	"开發piao":    InvalidTag,
	"开發票":       InvalidTag,
	"开發瞟":       InvalidTag,
	"开發漂":       InvalidTag,
	"开發镖":       InvalidTag,
	"开發標":       InvalidTag,
	"开發磦":       InvalidTag,
	"开發骠":       InvalidTag,
	"开發嘌":       InvalidTag,
	"开發缥":       InvalidTag,
	"开發螵":       InvalidTag,
	"开發彯":       InvalidTag,
	"开發慓":       InvalidTag,
	"开發剽":       InvalidTag,
	"开發徱":       InvalidTag,
	"开發僄":       InvalidTag,
	"开發飘":       InvalidTag,
	"开潑piao":    InvalidTag,
	"开潑票":       InvalidTag,
	"开潑瞟":       InvalidTag,
	"开潑漂":       InvalidTag,
	"开潑镖":       InvalidTag,
	"开潑標":       InvalidTag,
	"开潑磦":       InvalidTag,
	"开潑骠":       InvalidTag,
	"开潑嘌":       InvalidTag,
	"开潑缥":       InvalidTag,
	"开潑螵":       InvalidTag,
	"开潑彯":       InvalidTag,
	"开潑慓":       InvalidTag,
	"开潑剽":       InvalidTag,
	"开潑徱":       InvalidTag,
	"开潑僄":       InvalidTag,
	"开潑飘":       InvalidTag,
	"开泼piao":    InvalidTag,
	"开泼票":       InvalidTag,
	"开泼瞟":       InvalidTag,
	"开泼漂":       InvalidTag,
	"开泼镖":       InvalidTag,
	"开泼標":       InvalidTag,
	"开泼磦":       InvalidTag,
	"开泼骠":       InvalidTag,
	"开泼嘌":       InvalidTag,
	"开泼缥":       InvalidTag,
	"开泼螵":       InvalidTag,
	"开泼彯":       InvalidTag,
	"开泼慓":       InvalidTag,
	"开泼剽":       InvalidTag,
	"开泼徱":       InvalidTag,
	"开泼僄":       InvalidTag,
	"开泼飘":       InvalidTag,
	"开蕟piao":    InvalidTag,
	"开蕟票":       InvalidTag,
	"开蕟瞟":       InvalidTag,
	"开蕟漂":       InvalidTag,
	"开蕟镖":       InvalidTag,
	"开蕟標":       InvalidTag,
	"开蕟磦":       InvalidTag,
	"开蕟骠":       InvalidTag,
	"开蕟嘌":       InvalidTag,
	"开蕟缥":       InvalidTag,
	"开蕟螵":       InvalidTag,
	"开蕟彯":       InvalidTag,
	"开蕟慓":       InvalidTag,
	"开蕟剽":       InvalidTag,
	"开蕟徱":       InvalidTag,
	"开蕟僄":       InvalidTag,
	"开蕟飘":       InvalidTag,
	"开犮piao":    InvalidTag,
	"开犮票":       InvalidTag,
	"开犮瞟":       InvalidTag,
	"开犮漂":       InvalidTag,
	"开犮镖":       InvalidTag,
	"开犮標":       InvalidTag,
	"开犮磦":       InvalidTag,
	"开犮骠":       InvalidTag,
	"开犮嘌":       InvalidTag,
	"开犮缥":       InvalidTag,
	"开犮螵":       InvalidTag,
	"开犮彯":       InvalidTag,
	"开犮慓":       InvalidTag,
	"开犮剽":       InvalidTag,
	"开犮徱":       InvalidTag,
	"开犮僄":       InvalidTag,
	"开犮飘":       InvalidTag,
	"幵fapiao":   InvalidTag,
	"幵fa票":      InvalidTag,
	"幵fa瞟":      InvalidTag,
	"幵fa漂":      InvalidTag,
	"幵fa镖":      InvalidTag,
	"幵fa標":      InvalidTag,
	"幵fa磦":      InvalidTag,
	"幵fa骠":      InvalidTag,
	"幵fa嘌":      InvalidTag,
	"幵fa缥":      InvalidTag,
	"幵fa螵":      InvalidTag,
	"幵fa彯":      InvalidTag,
	"幵fa慓":      InvalidTag,
	"幵fa剽":      InvalidTag,
	"幵fa徱":      InvalidTag,
	"幵fa僄":      InvalidTag,
	"幵fa飘":      InvalidTag,
	"幵发piao":    InvalidTag,
	"幵发票":       InvalidTag,
	"幵发瞟":       InvalidTag,
	"幵发漂":       InvalidTag,
	"幵发镖":       InvalidTag,
	"幵发標":       InvalidTag,
	"幵发磦":       InvalidTag,
	"幵发骠":       InvalidTag,
	"幵发嘌":       InvalidTag,
	"幵发缥":       InvalidTag,
	"幵发螵":       InvalidTag,
	"幵发彯":       InvalidTag,
	"幵发慓":       InvalidTag,
	"幵发剽":       InvalidTag,
	"幵发徱":       InvalidTag,
	"幵发僄":       InvalidTag,
	"幵发飘":       InvalidTag,
	"幵發piao":    InvalidTag,
	"幵發票":       InvalidTag,
	"幵發瞟":       InvalidTag,
	"幵發漂":       InvalidTag,
	"幵發镖":       InvalidTag,
	"幵發標":       InvalidTag,
	"幵發磦":       InvalidTag,
	"幵發骠":       InvalidTag,
	"幵發嘌":       InvalidTag,
	"幵發缥":       InvalidTag,
	"幵發螵":       InvalidTag,
	"幵發彯":       InvalidTag,
	"幵發慓":       InvalidTag,
	"幵發剽":       InvalidTag,
	"幵發徱":       InvalidTag,
	"幵發僄":       InvalidTag,
	"幵發飘":       InvalidTag,
	"幵潑piao":    InvalidTag,
	"幵潑票":       InvalidTag,
	"幵潑瞟":       InvalidTag,
	"幵潑漂":       InvalidTag,
	"幵潑镖":       InvalidTag,
	"幵潑標":       InvalidTag,
	"幵潑磦":       InvalidTag,
	"幵潑骠":       InvalidTag,
	"幵潑嘌":       InvalidTag,
	"幵潑缥":       InvalidTag,
	"幵潑螵":       InvalidTag,
	"幵潑彯":       InvalidTag,
	"幵潑慓":       InvalidTag,
	"幵潑剽":       InvalidTag,
	"幵潑徱":       InvalidTag,
	"幵潑僄":       InvalidTag,
	"幵潑飘":       InvalidTag,
	"幵泼piao":    InvalidTag,
	"幵泼票":       InvalidTag,
	"幵泼瞟":       InvalidTag,
	"幵泼漂":       InvalidTag,
	"幵泼镖":       InvalidTag,
	"幵泼標":       InvalidTag,
	"幵泼磦":       InvalidTag,
	"幵泼骠":       InvalidTag,
	"幵泼嘌":       InvalidTag,
	"幵泼缥":       InvalidTag,
	"幵泼螵":       InvalidTag,
	"幵泼彯":       InvalidTag,
	"幵泼慓":       InvalidTag,
	"幵泼剽":       InvalidTag,
	"幵泼徱":       InvalidTag,
	"幵泼僄":       InvalidTag,
	"幵泼飘":       InvalidTag,
	"幵蕟piao":    InvalidTag,
	"幵蕟票":       InvalidTag,
	"幵蕟瞟":       InvalidTag,
	"幵蕟漂":       InvalidTag,
	"幵蕟镖":       InvalidTag,
	"幵蕟標":       InvalidTag,
	"幵蕟磦":       InvalidTag,
	"幵蕟骠":       InvalidTag,
	"幵蕟嘌":       InvalidTag,
	"幵蕟缥":       InvalidTag,
	"幵蕟螵":       InvalidTag,
	"幵蕟彯":       InvalidTag,
	"幵蕟慓":       InvalidTag,
	"幵蕟剽":       InvalidTag,
	"幵蕟徱":       InvalidTag,
	"幵蕟僄":       InvalidTag,
	"幵蕟飘":       InvalidTag,
	"幵犮piao":    InvalidTag,
	"幵犮票":       InvalidTag,
	"幵犮瞟":       InvalidTag,
	"幵犮漂":       InvalidTag,
	"幵犮镖":       InvalidTag,
	"幵犮標":       InvalidTag,
	"幵犮磦":       InvalidTag,
	"幵犮骠":       InvalidTag,
	"幵犮嘌":       InvalidTag,
	"幵犮缥":       InvalidTag,
	"幵犮螵":       InvalidTag,
	"幵犮彯":       InvalidTag,
	"幵犮慓":       InvalidTag,
	"幵犮剽":       InvalidTag,
	"幵犮徱":       InvalidTag,
	"幵犮僄":       InvalidTag,
	"幵犮飘":       InvalidTag,
	"開fapiao":   InvalidTag,
	"開fa票":      InvalidTag,
	"開fa瞟":      InvalidTag,
	"開fa漂":      InvalidTag,
	"開fa镖":      InvalidTag,
	"開fa標":      InvalidTag,
	"開fa磦":      InvalidTag,
	"開fa骠":      InvalidTag,
	"開fa嘌":      InvalidTag,
	"開fa缥":      InvalidTag,
	"開fa螵":      InvalidTag,
	"開fa彯":      InvalidTag,
	"開fa慓":      InvalidTag,
	"開fa剽":      InvalidTag,
	"開fa徱":      InvalidTag,
	"開fa僄":      InvalidTag,
	"開fa飘":      InvalidTag,
	"開发piao":    InvalidTag,
	"開发票":       InvalidTag,
	"開发瞟":       InvalidTag,
	"開发漂":       InvalidTag,
	"開发镖":       InvalidTag,
	"開发標":       InvalidTag,
	"開发磦":       InvalidTag,
	"開发骠":       InvalidTag,
	"開发嘌":       InvalidTag,
	"開发缥":       InvalidTag,
	"開发螵":       InvalidTag,
	"開发彯":       InvalidTag,
	"開发慓":       InvalidTag,
	"開发剽":       InvalidTag,
	"開发徱":       InvalidTag,
	"開发僄":       InvalidTag,
	"開发飘":       InvalidTag,
	"開發piao":    InvalidTag,
	"開發票":       InvalidTag,
	"開發瞟":       InvalidTag,
	"開發漂":       InvalidTag,
	"開發镖":       InvalidTag,
	"開發標":       InvalidTag,
	"開發磦":       InvalidTag,
	"開發骠":       InvalidTag,
	"開發嘌":       InvalidTag,
	"開發缥":       InvalidTag,
	"開發螵":       InvalidTag,
	"開發彯":       InvalidTag,
	"開發慓":       InvalidTag,
	"開發剽":       InvalidTag,
	"開發徱":       InvalidTag,
	"開發僄":       InvalidTag,
	"開發飘":       InvalidTag,
	"開潑piao":    InvalidTag,
	"開潑票":       InvalidTag,
	"開潑瞟":       InvalidTag,
	"開潑漂":       InvalidTag,
	"開潑镖":       InvalidTag,
	"開潑標":       InvalidTag,
	"開潑磦":       InvalidTag,
	"開潑骠":       InvalidTag,
	"開潑嘌":       InvalidTag,
	"開潑缥":       InvalidTag,
	"開潑螵":       InvalidTag,
	"開潑彯":       InvalidTag,
	"開潑慓":       InvalidTag,
	"開潑剽":       InvalidTag,
	"開潑徱":       InvalidTag,
	"開潑僄":       InvalidTag,
	"開潑飘":       InvalidTag,
	"開泼piao":    InvalidTag,
	"開泼票":       InvalidTag,
	"開泼瞟":       InvalidTag,
	"開泼漂":       InvalidTag,
	"開泼镖":       InvalidTag,
	"開泼標":       InvalidTag,
	"開泼磦":       InvalidTag,
	"開泼骠":       InvalidTag,
	"開泼嘌":       InvalidTag,
	"開泼缥":       InvalidTag,
	"開泼螵":       InvalidTag,
	"開泼彯":       InvalidTag,
	"開泼慓":       InvalidTag,
	"開泼剽":       InvalidTag,
	"開泼徱":       InvalidTag,
	"開泼僄":       InvalidTag,
	"開泼飘":       InvalidTag,
	"開蕟piao":    InvalidTag,
	"開蕟票":       InvalidTag,
	"開蕟瞟":       InvalidTag,
	"開蕟漂":       InvalidTag,
	"開蕟镖":       InvalidTag,
	"開蕟標":       InvalidTag,
	"開蕟磦":       InvalidTag,
	"開蕟骠":       InvalidTag,
	"開蕟嘌":       InvalidTag,
	"開蕟缥":       InvalidTag,
	"開蕟螵":       InvalidTag,
	"開蕟彯":       InvalidTag,
	"開蕟慓":       InvalidTag,
	"開蕟剽":       InvalidTag,
	"開蕟徱":       InvalidTag,
	"開蕟僄":       InvalidTag,
	"開蕟飘":       InvalidTag,
	"開犮piao":    InvalidTag,
	"開犮票":       InvalidTag,
	"開犮瞟":       InvalidTag,
	"開犮漂":       InvalidTag,
	"開犮镖":       InvalidTag,
	"開犮標":       InvalidTag,
	"開犮磦":       InvalidTag,
	"開犮骠":       InvalidTag,
	"開犮嘌":       InvalidTag,
	"開犮缥":       InvalidTag,
	"開犮螵":       InvalidTag,
	"開犮彯":       InvalidTag,
	"開犮慓":       InvalidTag,
	"開犮剽":       InvalidTag,
	"開犮徱":       InvalidTag,
	"開犮僄":       InvalidTag,
	"開犮飘":       InvalidTag,
	"闓fapiao":   InvalidTag,
	"闓fa票":      InvalidTag,
	"闓fa瞟":      InvalidTag,
	"闓fa漂":      InvalidTag,
	"闓fa镖":      InvalidTag,
	"闓fa標":      InvalidTag,
	"闓fa磦":      InvalidTag,
	"闓fa骠":      InvalidTag,
	"闓fa嘌":      InvalidTag,
	"闓fa缥":      InvalidTag,
	"闓fa螵":      InvalidTag,
	"闓fa彯":      InvalidTag,
	"闓fa慓":      InvalidTag,
	"闓fa剽":      InvalidTag,
	"闓fa徱":      InvalidTag,
	"闓fa僄":      InvalidTag,
	"闓fa飘":      InvalidTag,
	"闓发piao":    InvalidTag,
	"闓发票":       InvalidTag,
	"闓发瞟":       InvalidTag,
	"闓发漂":       InvalidTag,
	"闓发镖":       InvalidTag,
	"闓发標":       InvalidTag,
	"闓发磦":       InvalidTag,
	"闓发骠":       InvalidTag,
	"闓发嘌":       InvalidTag,
	"闓发缥":       InvalidTag,
	"闓发螵":       InvalidTag,
	"闓发彯":       InvalidTag,
	"闓发慓":       InvalidTag,
	"闓发剽":       InvalidTag,
	"闓发徱":       InvalidTag,
	"闓发僄":       InvalidTag,
	"闓发飘":       InvalidTag,
	"闓發piao":    InvalidTag,
	"闓發票":       InvalidTag,
	"闓發瞟":       InvalidTag,
	"闓發漂":       InvalidTag,
	"闓發镖":       InvalidTag,
	"闓發標":       InvalidTag,
	"闓發磦":       InvalidTag,
	"闓發骠":       InvalidTag,
	"闓發嘌":       InvalidTag,
	"闓發缥":       InvalidTag,
	"闓發螵":       InvalidTag,
	"闓發彯":       InvalidTag,
	"闓發慓":       InvalidTag,
	"闓發剽":       InvalidTag,
	"闓發徱":       InvalidTag,
	"闓發僄":       InvalidTag,
	"闓發飘":       InvalidTag,
	"闓潑piao":    InvalidTag,
	"闓潑票":       InvalidTag,
	"闓潑瞟":       InvalidTag,
	"闓潑漂":       InvalidTag,
	"闓潑镖":       InvalidTag,
	"闓潑標":       InvalidTag,
	"闓潑磦":       InvalidTag,
	"闓潑骠":       InvalidTag,
	"闓潑嘌":       InvalidTag,
	"闓潑缥":       InvalidTag,
	"闓潑螵":       InvalidTag,
	"闓潑彯":       InvalidTag,
	"闓潑慓":       InvalidTag,
	"闓潑剽":       InvalidTag,
	"闓潑徱":       InvalidTag,
	"闓潑僄":       InvalidTag,
	"闓潑飘":       InvalidTag,
	"闓泼piao":    InvalidTag,
	"闓泼票":       InvalidTag,
	"闓泼瞟":       InvalidTag,
	"闓泼漂":       InvalidTag,
	"闓泼镖":       InvalidTag,
	"闓泼標":       InvalidTag,
	"闓泼磦":       InvalidTag,
	"闓泼骠":       InvalidTag,
	"闓泼嘌":       InvalidTag,
	"闓泼缥":       InvalidTag,
	"闓泼螵":       InvalidTag,
	"闓泼彯":       InvalidTag,
	"闓泼慓":       InvalidTag,
	"闓泼剽":       InvalidTag,
	"闓泼徱":       InvalidTag,
	"闓泼僄":       InvalidTag,
	"闓泼飘":       InvalidTag,
	"闓蕟piao":    InvalidTag,
	"闓蕟票":       InvalidTag,
	"闓蕟瞟":       InvalidTag,
	"闓蕟漂":       InvalidTag,
	"闓蕟镖":       InvalidTag,
	"闓蕟標":       InvalidTag,
	"闓蕟磦":       InvalidTag,
	"闓蕟骠":       InvalidTag,
	"闓蕟嘌":       InvalidTag,
	"闓蕟缥":       InvalidTag,
	"闓蕟螵":       InvalidTag,
	"闓蕟彯":       InvalidTag,
	"闓蕟慓":       InvalidTag,
	"闓蕟剽":       InvalidTag,
	"闓蕟徱":       InvalidTag,
	"闓蕟僄":       InvalidTag,
	"闓蕟飘":       InvalidTag,
	"闓犮piao":    InvalidTag,
	"闓犮票":       InvalidTag,
	"闓犮瞟":       InvalidTag,
	"闓犮漂":       InvalidTag,
	"闓犮镖":       InvalidTag,
	"闓犮標":       InvalidTag,
	"闓犮磦":       InvalidTag,
	"闓犮骠":       InvalidTag,
	"闓犮嘌":       InvalidTag,
	"闓犮缥":       InvalidTag,
	"闓犮螵":       InvalidTag,
	"闓犮彯":       InvalidTag,
	"闓犮慓":       InvalidTag,
	"闓犮剽":       InvalidTag,
	"闓犮徱":       InvalidTag,
	"闓犮僄":       InvalidTag,
	"闓犮飘":       InvalidTag,
	"fapiao":    InvalidTag,
	"fa票":       InvalidTag,
	"fa瞟":       InvalidTag,
	"fa漂":       InvalidTag,
	"fa镖":       InvalidTag,
	"fa標":       InvalidTag,
	"fa磦":       InvalidTag,
	"fa骠":       InvalidTag,
	"fa嘌":       InvalidTag,
	"fa缥":       InvalidTag,
	"fa螵":       InvalidTag,
	"fa彯":       InvalidTag,
	"fa慓":       InvalidTag,
	"fa剽":       InvalidTag,
	"fa徱":       InvalidTag,
	"fa僄":       InvalidTag,
	"fa飘":       InvalidTag,
	"发piao":     InvalidTag,
	"发票":        UnknownTag,
	"发瞟":        InvalidTag,
	"发漂":        InvalidTag,
	"发镖":        InvalidTag,
	"发標":        InvalidTag,
	"发磦":        InvalidTag,
	"发骠":        InvalidTag,
	"发嘌":        InvalidTag,
	"发缥":        InvalidTag,
	"发螵":        InvalidTag,
	"发彯":        InvalidTag,
	"发慓":        InvalidTag,
	"发剽":        InvalidTag,
	"发徱":        InvalidTag,
	"发僄":        InvalidTag,
	"发飘":        InvalidTag,
	"發piao":     InvalidTag,
	"發票":        InvalidTag,
	"發瞟":        InvalidTag,
	"發漂":        InvalidTag,
	"發镖":        InvalidTag,
	"發標":        InvalidTag,
	"發磦":        InvalidTag,
	"發骠":        InvalidTag,
	"發嘌":        InvalidTag,
	"發缥":        InvalidTag,
	"發螵":        InvalidTag,
	"發彯":        InvalidTag,
	"發慓":        InvalidTag,
	"發剽":        InvalidTag,
	"發徱":        InvalidTag,
	"發僄":        InvalidTag,
	"發飘":        InvalidTag,
	"潑piao":     InvalidTag,
	"潑票":        InvalidTag,
	"潑瞟":        InvalidTag,
	"潑漂":        InvalidTag,
	"潑镖":        InvalidTag,
	"潑標":        InvalidTag,
	"潑磦":        InvalidTag,
	"潑骠":        InvalidTag,
	"潑嘌":        InvalidTag,
	"潑缥":        InvalidTag,
	"潑螵":        InvalidTag,
	"潑彯":        InvalidTag,
	"潑慓":        InvalidTag,
	"潑剽":        InvalidTag,
	"潑徱":        InvalidTag,
	"潑僄":        InvalidTag,
	"潑飘":        InvalidTag,
	"泼piao":     InvalidTag,
	"泼票":        InvalidTag,
	"泼瞟":        InvalidTag,
	"泼漂":        InvalidTag,
	"泼镖":        InvalidTag,
	"泼標":        InvalidTag,
	"泼磦":        InvalidTag,
	"泼骠":        InvalidTag,
	"泼嘌":        InvalidTag,
	"泼缥":        InvalidTag,
	"泼螵":        InvalidTag,
	"泼彯":        InvalidTag,
	"泼慓":        InvalidTag,
	"泼剽":        InvalidTag,
	"泼徱":        InvalidTag,
	"泼僄":        InvalidTag,
	"泼飘":        InvalidTag,
	"蕟piao":     InvalidTag,
	"蕟票":        InvalidTag,
	"蕟瞟":        InvalidTag,
	"蕟漂":        InvalidTag,
	"蕟镖":        InvalidTag,
	"蕟標":        InvalidTag,
	"蕟磦":        InvalidTag,
	"蕟骠":        InvalidTag,
	"蕟嘌":        InvalidTag,
	"蕟缥":        InvalidTag,
	"蕟螵":        InvalidTag,
	"蕟彯":        InvalidTag,
	"蕟慓":        InvalidTag,
	"蕟剽":        InvalidTag,
	"蕟徱":        InvalidTag,
	"蕟僄":        InvalidTag,
	"蕟飘":        InvalidTag,
	"犮piao":     InvalidTag,
	"犮票":        InvalidTag,
	"犮瞟":        InvalidTag,
	"犮漂":        InvalidTag,
	"犮镖":        InvalidTag,
	"犮標":        InvalidTag,
	"犮磦":        InvalidTag,
	"犮骠":        InvalidTag,
	"犮嘌":        InvalidTag,
	"犮缥":        InvalidTag,
	"犮螵":        InvalidTag,
	"犮彯":        InvalidTag,
	"犮慓":        InvalidTag,
	"犮剽":        InvalidTag,
	"犮徱":        InvalidTag,
	"犮僄":        InvalidTag,
	"犮飘":        InvalidTag,
	"daikai":    InvalidTag,
	"dai开":      InvalidTag,
	"dai幵":      InvalidTag,
	"dai開":      InvalidTag,
	"dai闓":      InvalidTag,
	"代kai":      InvalidTag,
	"代开":        UnknownTag,
	"代幵":        InvalidTag,
	"代開":        InvalidTag,
	"代闓":        InvalidTag,
	"岱kai":      InvalidTag,
	"岱开":        InvalidTag,
	"岱幵":        InvalidTag,
	"岱開":        InvalidTag,
	"岱闓":        InvalidTag,
	"玳kai":      InvalidTag,
	"玳开":        InvalidTag,
	"玳幵":        InvalidTag,
	"玳開":        InvalidTag,
	"玳闓":        InvalidTag,
	"带kai":      InvalidTag,
	"带开":        UnknownTag,
	"带幵":        InvalidTag,
	"带開":        InvalidTag,
	"带闓":        InvalidTag,
	"kaipiao":   InvalidTag,
	"kai票":      InvalidTag,
	"kai瞟":      InvalidTag,
	"kai漂":      InvalidTag,
	"kai镖":      InvalidTag,
	"kai標":      InvalidTag,
	"kai磦":      InvalidTag,
	"kai骠":      InvalidTag,
	"kai嘌":      InvalidTag,
	"kai缥":      InvalidTag,
	"kai螵":      InvalidTag,
	"kai彯":      InvalidTag,
	"kai慓":      InvalidTag,
	"kai剽":      InvalidTag,
	"kai徱":      InvalidTag,
	"kai僄":      InvalidTag,
	"kai飘":      InvalidTag,
	"开piao":     InvalidTag,
	"开票":        UnknownTag,
	"开瞟":        InvalidTag,
	"开漂":        InvalidTag,
	"开镖":        InvalidTag,
	"开標":        InvalidTag,
	"开磦":        InvalidTag,
	"开骠":        InvalidTag,
	"开嘌":        InvalidTag,
	"开缥":        InvalidTag,
	"开螵":        InvalidTag,
	"开彯":        InvalidTag,
	"开慓":        InvalidTag,
	"开剽":        InvalidTag,
	"开徱":        InvalidTag,
	"开僄":        InvalidTag,
	"开飘":        InvalidTag,
	"幵piao":     InvalidTag,
	"幵票":        InvalidTag,
	"幵瞟":        InvalidTag,
	"幵漂":        InvalidTag,
	"幵镖":        InvalidTag,
	"幵標":        InvalidTag,
	"幵磦":        InvalidTag,
	"幵骠":        InvalidTag,
	"幵嘌":        InvalidTag,
	"幵缥":        InvalidTag,
	"幵螵":        InvalidTag,
	"幵彯":        InvalidTag,
	"幵慓":        InvalidTag,
	"幵剽":        InvalidTag,
	"幵徱":        InvalidTag,
	"幵僄":        InvalidTag,
	"幵飘":        InvalidTag,
	"開piao":     InvalidTag,
	"開票":        InvalidTag,
	"開瞟":        InvalidTag,
	"開漂":        InvalidTag,
	"開镖":        InvalidTag,
	"開標":        InvalidTag,
	"開磦":        InvalidTag,
	"開骠":        InvalidTag,
	"開嘌":        InvalidTag,
	"開缥":        InvalidTag,
	"開螵":        InvalidTag,
	"開彯":        InvalidTag,
	"開慓":        InvalidTag,
	"開剽":        InvalidTag,
	"開徱":        InvalidTag,
	"開僄":        InvalidTag,
	"開飘":        InvalidTag,
	"闓piao":     InvalidTag,
	"闓票":        InvalidTag,
	"闓瞟":        InvalidTag,
	"闓漂":        InvalidTag,
	"闓镖":        InvalidTag,
	"闓標":        InvalidTag,
	"闓磦":        InvalidTag,
	"闓骠":        InvalidTag,
	"闓嘌":        InvalidTag,
	"闓缥":        InvalidTag,
	"闓螵":        InvalidTag,
	"闓彯":        InvalidTag,
	"闓慓":        InvalidTag,
	"闓剽":        InvalidTag,
	"闓徱":        InvalidTag,
	"闓僄":        InvalidTag,
	"闓飘":        InvalidTag,
	"税piao":     InvalidTag,
	"税票":        UnknownTag,
	"税瞟":        InvalidTag,
	"税漂":        InvalidTag,
	"税镖":        InvalidTag,
	"税標":        InvalidTag,
	"税磦":        InvalidTag,
	"税骠":        InvalidTag,
	"税嘌":        InvalidTag,
	"税缥":        InvalidTag,
	"税螵":        InvalidTag,
	"税彯":        InvalidTag,
	"税慓":        InvalidTag,
	"税剽":        InvalidTag,
	"税徱":        InvalidTag,
	"税僄":        InvalidTag,
	"税飘":        InvalidTag,
	"piao据":     InvalidTag,
	"票据":        UnknownTag,
	"瞟据":        InvalidTag,
	"漂据":        InvalidTag,
	"镖据":        InvalidTag,
	"標据":        InvalidTag,
	"磦据":        InvalidTag,
	"骠据":        InvalidTag,
	"嘌据":        InvalidTag,
	"缥据":        InvalidTag,
	"螵据":        InvalidTag,
	"彯据":        InvalidTag,
	"慓据":        InvalidTag,
	"剽据":        InvalidTag,
	"徱据":        InvalidTag,
	"僄据":        InvalidTag,
	"飘据":        InvalidTag,
	"zengpiao":  InvalidTag,
	"zeng票":     InvalidTag,
	"zeng瞟":     InvalidTag,
	"zeng漂":     InvalidTag,
	"zeng镖":     InvalidTag,
	"zeng標":     InvalidTag,
	"zeng磦":     InvalidTag,
	"zeng骠":     InvalidTag,
	"zeng嘌":     InvalidTag,
	"zeng缥":     InvalidTag,
	"zeng螵":     InvalidTag,
	"zeng彯":     InvalidTag,
	"zeng慓":     InvalidTag,
	"zeng剽":     InvalidTag,
	"zeng徱":     InvalidTag,
	"zeng僄":     InvalidTag,
	"zeng飘":     InvalidTag,
	"增piao":     InvalidTag,
	"增票":        UnknownTag,
	"增瞟":        InvalidTag,
	"增漂":        InvalidTag,
	"增镖":        InvalidTag,
	"增標":        InvalidTag,
	"增磦":        InvalidTag,
	"增骠":        InvalidTag,
	"增嘌":        InvalidTag,
	"增缥":        InvalidTag,
	"增螵":        InvalidTag,
	"增彯":        InvalidTag,
	"增慓":        InvalidTag,
	"增剽":        InvalidTag,
	"增徱":        InvalidTag,
	"增僄":        InvalidTag,
	"增飘":        InvalidTag,
	"磳piao":     InvalidTag,
	"磳票":        InvalidTag,
	"磳瞟":        InvalidTag,
	"磳漂":        InvalidTag,
	"磳镖":        InvalidTag,
	"磳標":        InvalidTag,
	"磳磦":        InvalidTag,
	"磳骠":        InvalidTag,
	"磳嘌":        InvalidTag,
	"磳缥":        InvalidTag,
	"磳螵":        InvalidTag,
	"磳彯":        InvalidTag,
	"磳慓":        InvalidTag,
	"磳剽":        InvalidTag,
	"磳徱":        InvalidTag,
	"磳僄":        InvalidTag,
	"磳飘":        InvalidTag,
	"zengzhi票":  InvalidTag,
	"zeng值票":    InvalidTag,
	"zeng徝票":    InvalidTag,
	"zeng値票":    InvalidTag,
	"增zhi票":     InvalidTag,
	"增值票":       UnknownTag,
	"增徝票":       InvalidTag,
	"增値票":       InvalidTag,
	"磳zhi票":     InvalidTag,
	"磳值票":       InvalidTag,
	"磳徝票":       InvalidTag,
	"磳値票":       InvalidTag,
	"増zhi票":     InvalidTag,
	"増值票":       InvalidTag,
	"増徝票":       InvalidTag,
	"増値票":       InvalidTag,
	"禾兑":        InvalidTag,
	"真票":        InvalidTag,
	"稅piao":     InvalidTag,
	"稅票":        InvalidTag,
	"稅瞟":        InvalidTag,
	"稅漂":        InvalidTag,
	"稅镖":        InvalidTag,
	"稅標":        InvalidTag,
	"稅磦":        InvalidTag,
	"稅骠":        InvalidTag,
	"稅嘌":        InvalidTag,
	"稅缥":        InvalidTag,
	"稅螵":        InvalidTag,
	"稅彯":        InvalidTag,
	"稅慓":        InvalidTag,
	"稅剽":        InvalidTag,
	"稅徱":        InvalidTag,
	"稅僄":        InvalidTag,
	"稅飘":        InvalidTag,
	"办理全国业务税票":  InvalidTag,
	"请问需要开发票":   InvalidTag,
	"需要发票开吗":    InvalidTag,
	//定义广告类的关键字
	"折优惠":   InvalidTag,
	"开业大酬宾": InvalidTag,
	"充值优惠":  InvalidTag,
}
