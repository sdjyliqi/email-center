package main

import (
	"email-center/reader"
	"email-center/tryagain"
	"email-center/utils"
	"strings"
)

func pickupFiles(fileNames string) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	files := strings.Split(fileNames, ";")
	for _, v := range files {
		filePickup, err := pickupOnFile(v)
		if err != nil {
			return nil, err
		}
		result[v] = filePickup
	}
	return result, nil
}

func pickupOnFile(f string) (map[string]int, error) {
	content, err := reader.FReader.ReadTxt(f)
	if err != nil {
		return nil, err
	}
	//识别内容中的身份证号
	strContent := string(content)
	cardIDs := tryagain.ExtractCardIDs(strContent)
	phoneIDs, _ := utils.ExtractMobilePhone(strContent)
	addrInfos := tryagain.PickAddr(strContent)
	fixPhoneIDs := tryagain.PickTelephone(strContent)
	bankIDs := tryagain.PickBankCard(strContent)
	result := map[string]int{
		"card":  len(cardIDs),
		"phone": len(phoneIDs),
		"addr":  len(addrInfos),
		"fix":   len(fixPhoneIDs),
		"bank":  len(bankIDs),
	}
	return result, nil
}
