package utils

import "fmt"

//CreateBlackIndex ... 创建发票类索引关键字
func CreateFAPIAOBlackIndex() []string {
	var emailKeyWords = []string{"开具发票"}
	//代开发票三个字的组合
	var dai = []string{"dai", "代", "岱", "玳", "带"}
	var kai = []string{"kai", "开", "幵", "開", "闓"}
	var fa = []string{"fa", "发", "發", "潑", "泼", "蕟", "犮"}
	var piao = []string{"piao", "票", "瞟", "漂", "镖", "標", "磦", "骠", "嘌", "缥", "螵", "彯", "慓", "剽", "徱", "僄", "飘"}
	//增值二字的组合
	var zeng = []string{"zeng", "增", "磳"}
	var zhi = []string{"zhi", "值", "徝", "値"}
	//存放开发票三个字排列组合结果
	for _, v1 := range kai {
		for _, v2 := range fa {
			for _, v3 := range piao {
				fullName := fmt.Sprintf("%s%s%s", v1, v2, v3)
				emailKeyWords = append(emailKeyWords, fullName)
			}
		}
	}
	//存放发、票二字排列组合结果，有了“发票”还存“开发票”的原因：“开”字可以直接判断发票类的真、假。
	for _, v1 := range fa {
		for _, v2 := range piao {
			fullName := fmt.Sprintf("%s%s", v1, v2)
			emailKeyWords = append(emailKeyWords, fullName)
		}
	}
	//存放代、开二字排列组合结果
	for _, v1 := range dai {
		for _, v2 := range kai {
			fullName := fmt.Sprintf("%s%s", v1, v2)
			emailKeyWords = append(emailKeyWords, fullName)
		}
	}
	//存放开、票二字排列组合结果
	for _, v1 := range kai {
		for _, v2 := range piao {
			fullName := fmt.Sprintf("%s%s", v1, v2)
			emailKeyWords = append(emailKeyWords, fullName)
		}
	}

	//税票一词的组合
	for _, v1 := range piao {
		fullName := fmt.Sprintf("税%s", v1)
		emailKeyWords = append(emailKeyWords, fullName)
	}
	//票据一词的组合
	for _, v1 := range piao {
		fullName := fmt.Sprintf("%s据", v1)
		emailKeyWords = append(emailKeyWords, fullName)
	}
	//增票一词的组合
	for _, v1 := range zeng {
		for _, v2 := range piao {
			fullName := fmt.Sprintf("%s%s", v1, v2)
			emailKeyWords = append(emailKeyWords, fullName)
		}
	}
	//增值票三字组合
	for _, v1 := range zeng {
		for _, v2 := range zhi {
			fullName := fmt.Sprintf("%s%s票", v1, v2)
			emailKeyWords = append(emailKeyWords, fullName)
		}
	}
	//需特殊处理的发票类判别字符集
	var special = []string{"禾兑", "真票"}
	emailKeyWords = append(emailKeyWords, special...)
	return emailKeyWords
}
