package utils

//init() ...完成工具类相关的初始化相关事项
func init() {
	InitURLDomainAC()     //初始化AC自动机
	InitCategoryWordsAC() //初始化分类的AC自动机
	InitMySQL("root:Bit0123456789!@tcp(114.55.139.105:3306)/email-center?charset=utf8mb4", true)
}
