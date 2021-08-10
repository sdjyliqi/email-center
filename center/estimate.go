package center

import (
	"email-center/model"
	"email-center/utils"
	"fmt"
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
	var amendChars []rune
	chars := []rune(content)
	for _, v := range chars {
		if v < 'A' || v > 'z' && v <= 255 {
			continue
		}
		amendChars = append(amendChars, v)
	}
	newSubject := string(amendChars)
	for _, v := range e.assistCharacter {
		newSubject = strings.ReplaceAll(newSubject, v, "")
	}
	return newSubject
}

//GetCategory ...获取待鉴别邮件的分类
func (e estimate) GetCategory(content string) (utils.Category, string) {
	newSubject := e.AmendSubject(content)
	fmt.Println("====subject after amend:", content, newSubject)
	return utils.GetCategoryIdx(newSubject)
}

//GetCategory ...获取待鉴别邮件的分类
func (e estimate) AuditAllEmailItems() error {
	items, err := model.BodyModel.GetAllItems()
	if err != nil {
		return nil
	}
	for _, v := range items {
		//先计算其分类，然后更新到数据库中，后续可以比较了存入数据的分类是否和计算的分类一致。
		partition, tag := e.GetCategory(v.Subject)
		v.Partition = partition.Name()
		err = model.BodyModel.UpdateItemCols(v, []string{"partition"})
		if err != nil {
			return err
		}
		//根据提取出的数据判断真假,首先通过预先定义的标签映射表，判断是否非法的类型
		fmt.Println(tag)
		val, ok := utils.TagProperty[tag]
		if ok && val == utils.InvalidTag {

			continue
		}

	}
	return nil
}
