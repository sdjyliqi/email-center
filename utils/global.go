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
	DirtyCategory   Category = 3 //色情类
)

func (c Category) Name() string {
	switch c {
	case BillCategory:
		return "发票类"
	case AdvertCategory:
		return "广告类"
	case DirtyCategory:
		return "色情类"
	default:
		return "未知分类"
	}
}

var WebFormat = "(http|https)://[a-z0-9\\.]+"
var shortWebFormat = "[a-z0-9\\.]{2,12}.(cn|com)"

var PhoneFormat = "(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}[^@A-Za-z]" //手机号码格式
var TimeFormat = "2006-01-02 15:04:05"
var SMTP163Host = "smtp.163.com:25"
var ADBlackWords = []string{"微信", "vx", "qq", "扣扣"}

//定义广告类分类关键字
var TagADProperty = map[string]LegalTag{
	"基金":     UnknownTag,
	"证券":     UnknownTag,
	"期货":     UnknownTag,
	"股票":     UnknownTag,
	"资产管理":   UnknownTag,
	"托管资产":   UnknownTag,
	"保险":     UnknownTag,
	"一条龙服务":  UnknownTag,
	"限时免费":   InvalidTag,
	"直播报名":   InvalidTag,
	"可定制内训":  InvalidTag,
	"内训请详询":  InvalidTag,
	"開催":     InvalidTag,
	"代办发票":   InvalidTag,
	"开各地正规":  InvalidTag,
	"公開":     InvalidTag,
	"優惠":     InvalidTag,
	"优惠":     InvalidTag,
	"元领":     InvalidTag,
	"開始":     InvalidTag,
	"疯抢":     InvalidTag,
	"1折":     InvalidTag,
	"2折":     InvalidTag,
	"3折":     InvalidTag,
	"4折":     InvalidTag,
	"5折":     InvalidTag,
	"6折":     InvalidTag,
	"7折":     InvalidTag,
	"8折":     InvalidTag,
	"9折":     InvalidTag,
	"一折":     InvalidTag,
	"二折":     InvalidTag,
	"三折":     InvalidTag,
	"四折":     InvalidTag,
	"五折":     InvalidTag,
	"六折":     InvalidTag,
	"七折":     InvalidTag,
	"八折":     InvalidTag,
	"九折":     InvalidTag,
	"会员卡":    InvalidTag,
	"充值送礼":   InvalidTag,
	"优惠券":    InvalidTag,
	"大酬宾":    InvalidTag,
	"新店开业":   InvalidTag,
	"免费送":    InvalidTag,
	"折优惠":    InvalidTag,
	"每日信用管家": ValidTag,
	"报名":     UnknownTag,
	"注册":     UnknownTag,
	"培训":     UnknownTag,
	"内训":     UnknownTag,
	"训练":     UnknownTag,
}

//定义发票分类的关键字，通过关键字可能会判断出是否为异常短信
var TagBillProperty = TagPropertyDict
