package ac

import (
	"email-center/model"
	"email-center/utils"
	"github.com/gansidui/ahocorasick"
	"sort"
	"strings"
)

//ashjashjas
var DomainACMatch *ahocorasick.Matcher
var CategoryACMatch *ahocorasick.Matcher
var HighlightsACMatch *ahocorasick.Matcher
var CustomerServiceACMatch *ahocorasick.Matcher
var ADBlackWordACMatch *ahocorasick.Matcher

var URLDomains = []string{"jd.com", "dangdang.com", "cebbank.com", "suning.com"}

//billCategoryWords ...广告相关分类的关键字
var billCategoryWords = []string{}
var advsCategoryWords = []string{}
var categoryBox = map[string]utils.Category{}
var AllCategoryWords = []string{}
var HighlightsWords = []string{}
var customerServiceWords = []string{}

//InitURLDomainAC ...初始化AC自动机
//todo  如果不用直接删除
func InitURLDomainAC() {
	DomainACMatch = ahocorasick.NewMatcher()
	DomainACMatch.Build(URLDomains)
}

//GetCategoryIdx ...获取邮件的分类索引名称,和索引词
func GetCategoryIdx(idx string) (utils.Category, string) {
	var tags utils.StringSlice
	idxList := CategoryACMatch.Match(idx)
	for _, v := range idxList {
		tag := AllCategoryWords[v]
		tags = append(tags, tag)
	}
	sort.Sort(tags)
	for _, v := range []string(tags) {
		result, ok := categoryBox[v]
		if ok {
			return result, v
		}
	}
	return utils.UnknownCategory, ""
}

//InitCategoryWordsAC ...构建分类的自动机
func InitCategoryWordsAC() {
	//初始化所有的分类关键字列表
	for _, v := range billCategoryWords {
		categoryBox[v] = utils.BillCategory
	}

	for _, v := range advsCategoryWords {
		categoryBox[v] = utils.AdvertCategory
	}
	AllCategoryWords = append(billCategoryWords, advsCategoryWords...)
	CategoryACMatch = ahocorasick.NewMatcher()
	CategoryACMatch.Build(AllCategoryWords)
}

//InitHighlightsAC ...构建分类的自动机
func InitHighlightsAC() {
	//初始化所有的分类关键字列表
	items, _ := model.DomainModel.GetAllItems()
	for _, v := range items {
		if len(v.Hotline) > 1 {
			words := strings.Split(v.Hotline, ",")
			HighlightsWords = append(HighlightsWords, words...)
		}
		if len(v.Highlights) > 1 {
			words := strings.Split(v.Highlights, ",")
			HighlightsWords = append(HighlightsWords, words...)
		}
	}
	HighlightsACMatch = ahocorasick.NewMatcher()
	HighlightsACMatch.Build(HighlightsWords)
}

func GetWhiteHighlights(idx string) []string {
	var words []string
	idxList := HighlightsACMatch.Match(idx)
	for _, v := range idxList {
		tag := HighlightsWords[v]
		words = append(words, tag)
	}
	return words
}

//InitCustomerServiceAC ...构建官方客服电话的自动机
func InitCustomerServiceAC() {
	//初始化所有的官方客服电话关键字列表
	items, _ := model.DomainModel.GetAllItems()
	for _, v := range items {
		if len(v.Hotline) > 1 {
			ids := strings.Split(v.Hotline, ",")
			customerServiceWords = append(customerServiceWords, ids...)
		}
	}
	CustomerServiceACMatch = ahocorasick.NewMatcher()
	CustomerServiceACMatch.Build(customerServiceWords)
}

//InitADBlackWordsServiceAC ...构建广告类所属的黑名单词
func InitADBlackWordsServiceAC() {
	ADBlackWordACMatch = ahocorasick.NewMatcher()
	ADBlackWordACMatch.Build(utils.ADBlackWords)
}

//GetCustomerServiceIDs ... 利用AC自动机获取官方客服电话
func GetCustomerServiceIDs(content string) []string {
	var words []string
	idxList := CustomerServiceACMatch.Match(content)
	for _, v := range idxList {
		tag := customerServiceWords[v]
		words = append(words, tag)
	}
	return words
}

//GetADBlackWords ... 利用AC获取黑名单词
func GetADBlackWords(content string) []string {
	var words []string
	idxList := ADBlackWordACMatch.Match(content)
	for _, v := range idxList {
		tag := utils.ADBlackWords[v]
		words = append(words, tag)
	}
	return words
}
