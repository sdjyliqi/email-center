package utils

import (
	"github.com/gansidui/ahocorasick"
	"sort"
)

//ashjashjas
var DomainACMatch *ahocorasick.Matcher
var CategoryACMatch *ahocorasick.Matcher

var URLDomains = []string{"jd.com", "dangdang.com", "cebbank.com", "suning.com"}

//billCategoryWords ...广告相关分类的关键字
var billCategoryWords = []string{}
var advsCategoryWords = []string{"充值送礼", "优惠券"}
var categoryBox = map[string]Category{}
var AllCategoryWords = []string{}

//InitURLDomainAC ...初始化AC自动机
func InitURLDomainAC() {
	DomainACMatch = ahocorasick.NewMatcher()
	DomainACMatch.Build(URLDomains)
}

//GetCategoryIdx ...获取邮件的分类索引名称,和索引词
func GetCategoryIdx(idx string) (Category, string) {
	var tags StringSlice
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
	return UnknownCategory, ""
}

//InitCategoryWordsAC ...构建分类的自动机
func InitCategoryWordsAC() {
	//初始化所有的分类关键字列表
	for _, v := range billCategoryWords {
		categoryBox[v] = BillCategory
	}

	for _, v := range advsCategoryWords {
		categoryBox[v] = AdvertCategory
	}
	AllCategoryWords = append(billCategoryWords, advsCategoryWords...)
	CategoryACMatch = ahocorasick.NewMatcher()
	CategoryACMatch.Build(AllCategoryWords)
}
