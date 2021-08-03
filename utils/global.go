package utils

//定义邮件合法性标记
type LegalTag int8

const (
	UnknownTag     LegalTag = 0  //无状态
	ValidTag       LegalTag = 1  //合法
	InvalidTag     LegalTag = 2  //非法
	lenMobilePhone int      = 11 //移动电话标准长度
)

var WebFormat = "(http|https)://[a-z0-9\\.]+"
var shortWebFormat = "[a-z0-9\\.]{2,12}.(cn|com)"

var PhoneFormat = "(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}" //手机号码格式

var TICKETBlackIdx = []string{
	"开具发票", "kaifapiao", "kaifa票", "kaifa瞟", "kaifa漂", "kaifa镖",
	"kaifa標", "kai发piao", "kai发票", "kai发瞟", "kai发漂", "kai发镖", "kai发標",
	"kai發piao", "kai發票", "kai發瞟", "kai發漂", "kai發镖", "kai發標", "kai潑piao",
	"kai潑票", "kai潑瞟", "kai潑漂", "kai潑镖", "kai潑標", "开fapiao", "开fa票", "开fa瞟", "开fa漂",
	"开fa镖", "开fa標", "开发piao", "开发票", "开发瞟", "开发漂", "开发镖", "开发標", "开發piao", "开發票", "开發瞟", "开發漂", "开發镖", "开發標", "开潑piao", "开潑票", "开潑瞟", "开潑漂", "开潑镖", "开潑標", "幵fapiao", "幵fa票", "幵fa瞟", "幵fa漂", "幵fa镖", "幵fa標", "幵发piao", "幵发票", "幵发瞟", "幵发漂", "幵发镖", "幵发標", "幵發piao", "幵發票", "幵發瞟", "幵發漂", "幵發镖", "幵發標", "幵潑piao", "幵潑票", "幵潑瞟", "幵潑漂", "幵潑镖", "幵潑標", "開fapiao", "開fa票", "開fa瞟", "開fa漂", "開fa镖", "開fa標", "開发piao", "開发票", "開发瞟", "開发漂", "開发镖", "開发標", "開發piao", "開發票", "開發瞟", "開發漂", "開發镖", "開發標", "開潑piao",
	"開潑票", "開潑瞟", "開潑漂", "開潑镖", "開潑標", "闓fapiao", "闓fa票", "闓fa瞟",
	"闓fa漂", "闓fa镖", "闓fa標", "闓发piao", "闓发票", "闓发瞟", "闓发漂", "闓发镖",
	"闓发標", "闓發piao", "闓發票", "闓發瞟", "闓發漂", "闓發镖", "闓發標", "闓潑piao",
	"闓潑票", "闓潑瞟", "闓潑漂", "闓潑镖", "闓潑標",
}
