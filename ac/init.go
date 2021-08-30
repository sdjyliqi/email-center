package ac

import "email-center/utils"

//init() ...完成工具类相关的初始化相关事项
func init() {
	//初始化一下分类的相关词汇
	for k, _ := range utils.TagProperty {
		billCategoryWords = append(billCategoryWords, k)
	}
	InitURLDomainAC()       //初始化AC自动机
	InitCategoryWordsAC()   //初始化分类的AC自动机
	InitHighlightsAC()      //初始化白名单关键字，主要包括大型公司的客服电话或者关键字如JD.com
	InitCustomerServiceAC() //构建所有客服电话的自动机
}
