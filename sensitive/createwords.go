package main

import (
	"email-center/model"
	"fmt"
)

//[]*portion的结构体中提取word列参数并存入[]string中
func portionToString(idx string) []string {
	var items []string
	t, err := model.PortionModel.GetIdxItems(idx)
	if err != nil {
		return nil
	}
	for _, v := range t {
		items = append(items, v.Word)
	}
	return items
}

func CreateKeywords() []string {
	//开发票、发票、代开、开票、税票、增票、票据、增值票
	var keywords = []string{}
	var kai = portionToString("kai")
	var dai = portionToString("dai")
	var fa = portionToString("fa")
	var piao = portionToString("piao")
	var zeng = portionToString("zeng")
	var zhi = portionToString("zhi")
	var shui = portionToString("shui")
	var ju = portionToString("ju")

	for _, v1 := range dai {
		for _, v2 := range kai {
			for _, v3 := range fa {
				for _, v4 := range piao {
					fullName := fmt.Sprintf("%s%s%s%s", v1, v2, v3, v4)
					keywords = append(keywords, fullName)
				}
			}
		}
	}

	for _, v1 := range kai {
		for _, v2 := range fa {
			for _, v3 := range piao {
				fullName := fmt.Sprintf("%s%s%s", v1, v2, v3)
				keywords = append(keywords, fullName)
			}
		}
	}

	for _, v1 := range fa {
		for _, v2 := range piao {
			fullName := fmt.Sprintf("%s%s", v1, v2)
			keywords = append(keywords, fullName)
		}
	}

	for _, v1 := range dai {
		for _, v2 := range kai {
			fullName := fmt.Sprintf("%s%s", v1, v2)
			keywords = append(keywords, fullName)
		}
	}

	for _, v1 := range kai {
		for _, v2 := range piao {
			fullName := fmt.Sprintf("%s%s", v1, v2)
			keywords = append(keywords, fullName)
		}
	}

	for _, v1 := range shui {
		for _, v2 := range piao {
			fullName := fmt.Sprintf("%s%s", v1, v2)
			keywords = append(keywords, fullName)
		}
	}

	for _, v1 := range zeng {
		for _, v2 := range piao {
			fullName := fmt.Sprintf("%s%s", v1, v2)
			keywords = append(keywords, fullName)
		}
	}

	for _, v1 := range zeng {
		for _, v2 := range zhi {
			for _, v3 := range piao {
				fullName := fmt.Sprintf("%s%s%s", v1, v2, v3)
				keywords = append(keywords, fullName)
			}
		}
	}

	for _, v1 := range piao {
		for _, v2 := range ju {
			fullName := fmt.Sprintf("%s%s", v1, v2)
			keywords = append(keywords, fullName)
		}
	}

	items, err := model.SensitiveModel.GetAllItems()
	if err != nil {
		return nil
	}
	for _, v := range items {
		keywords = append(keywords, v.Words)
	}

	return keywords
}

func createMap() map[string]string {
	dict := map[string]string{}
	key := CreateKeywords()
	for _, v := range key {
		dict[v] = "InvalidTag"
	}
	//unknowtag的名单
	var unknow = []string{"增值票", "票据", "税票", "开票", "带开", "发票", "开发票"}
	for _, v := range unknow {
		dict[v] = "UnknownTag"
	}
	return dict
}
