package utils

import "fmt"

//CreateBlackIndex ... 创建发票类索引关键字
func CreateFAPIAOBlackIndex() []string {
	var blackWords = []string{"开具发票"}
	//开发票三个字的组合
	var kai = []string{"kai", "开", "幵", "開", "闓"}
	var fa = []string{"fa", "发", "發", "潑"}
	var piao = []string{"piao", "票", "瞟", "漂", "镖", "標"}
	//存放开发票三个字排列组合结果
	for _, v1 := range kai {
		for _, v2 := range fa {
			for _, v3 := range piao {
				fullName := fmt.Sprintf("%s%s%s", v1, v2, v3)
				blackWords = append(blackWords, fullName)
			}
		}
	}
	return blackWords
}
