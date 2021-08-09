package center

import (
	"email-center/model"
	"email-center/utils"
	"strings"
)

//判断维度是否是否为真假

type estimate struct {
	assistCharacter []string //设置
	senderDomains   map[string]utils.LegalTag
}

//CreateEstimate 创建一个鉴定实例
func CreateEstimate() (*estimate, error) {
	var characters []string
	domains := map[string]utils.LegalTag{}
	items, err := model.AssistCharacter{}.GetAllItems()
	if err != nil {
		return nil, err
	}
	for _, v := range items {
		characters = append(characters, v.Character)
	}
	domainItems, err := model.Domain{}.GetAllItems()
	if err != nil {
		return nil, err
	}
	//目前数据均是白名单，即合法数据
	for _, v := range domainItems {
		domains[v.SenderName] = utils.ValidTag
	}
	return &estimate{assistCharacter: characters, senderDomains: domains}, nil
}

//AmendSubject ...修正标题,剔除一些无用的字符
func (e estimate) AmendSubject(content string) string {
	amendSubject := content
	for _, v := range e.assistCharacter {
		amendSubject = strings.ReplaceAll(content, v, "")
	}
	return amendSubject
}

//GetCategory ...获取待鉴别邮件的分类
func (e estimate) GetCategory(content string) string {
	content = e.AmendSubject(content)
	return utils.GetCategoryIdx(content).Name()
}
