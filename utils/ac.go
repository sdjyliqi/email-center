package utils

import (
	"fmt"
	"github.com/gansidui/ahocorasick"
)

//ashjashjas
var DomainACMatch *ahocorasick.Matcher
var CategoryACMatch *ahocorasick.Matcher

var URLDomains = []string{"jd.com", "dangdang.com", "cebbank.com", "suning.com"}
var billCategoryWords = []string{"开发票", "開发票", "kai发票", "開發票"}
var advsCategoryWords = []string{"充值送礼", "优惠券"}
var categoryBox = map[string]Category{}
var AllCategoryWords = []string{}

//InitURLDomainAC ...初始化AC自动机
func InitURLDomainAC() {
	fmt.Println("====InitURLDomainAC======初始化AC自动机==========")
	DomainACMatch = ahocorasick.NewMatcher()
	DomainACMatch.Build(URLDomains)
}

//GetCategoryIdx ...获取邮件的分类索引名称
func GetCategoryIdx(idx string) Category {
	result := CategoryACMatch.Match(idx)
	for _, v := range result {
		tag := AllCategoryWords[v]
		result, ok := categoryBox[tag]
		if ok {
			return result
		}
	}
	return UnknownCategory
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
	fmt.Println("====InitCategoryWordsAC====初始化AC自动机==========")
	CategoryACMatch = ahocorasick.NewMatcher()
	CategoryACMatch.Build(AllCategoryWords)
}
