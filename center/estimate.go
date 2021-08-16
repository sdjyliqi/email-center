package center

import (
	"email-center/ac"
	"email-center/model"
	"email-center/utils"
	"fmt"
	"strings"
)

//判断维度是否是否为真假
type estimate struct {
	assistCharacter []string //设置
	amendCharacters []*model.Amend
	domainWhite     map[string]utils.LegalTag
}

// CreateEstimate ... 创建一个鉴定实例
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
	//初始化域名白名单，即合法数据
	for _, v := range domainItems {
		domains[v.Official] = utils.ValidTag
	}
	//
	amendItems, err := model.AmendModel.GetAllItems()
	if err != nil {
		return nil, err
	}

	return &estimate{
		assistCharacter: characters,
		amendCharacters: amendItems,
		domainWhite:     map[string]utils.LegalTag{},
	}, nil
}

//AmendSubject ...修正标题,剔除一些无用的字符
func (e estimate) AmendSubjectForCategory(content string) string {
	var amendChars []rune
	chars := []rune(content)
	for _, v := range chars {
		if v < 'A' || v > 'Z' && v < 'a' || v > 'z' && v <= 255 {
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

//AmendSubject ...修正标题,剔除一些无用的字符
func (e estimate) AmendSubject(content string) string {
	var amendChars []rune
	chars := []rune(content)
	for _, v := range chars {
		if v < 'A' || v > 'Z' && v < 'a' || v > 'z' && v <= 255 {
			continue
		}
		amendChars = append(amendChars, v)
	}
	newSubject := string(amendChars)
	for _, v := range e.amendCharacters {
		newSubject = strings.ReplaceAll(newSubject, v.Raw, v.Replace)
	}
	return newSubject
}
func (e estimate) AmendSubjectExtent(content string) string {
	newSubject := content
	for _, v := range e.amendCharacters {
		newSubject = strings.ReplaceAll(newSubject, v.Raw, v.Replace)
	}
	return newSubject

}

//AmendBody ...修正邮件正文
func (e estimate) AmendBody(content string) string {
	var amendChars []rune
	chars := []rune(content)
	for _, v := range chars {
		if v < '0' || v > '9' && v < 'A' || v > 'Z' && v < 'a' || v > 'z' && v <= 255 {
			continue
		}
		amendChars = append(amendChars, v)
	}
	newContent := string(amendChars)
	for _, v := range e.amendCharacters {
		newContent = strings.ReplaceAll(newContent, v.Raw, v.Replace)
	}
	for _, v := range e.assistCharacter {
		newContent = strings.ReplaceAll(newContent, v, "")
	}
	return newContent
}

//GetCategory ...获取待鉴别邮件的分类
func (e estimate) GetCategory(subject, body string) (utils.Category, string) {
	//如果通过标题无法判断出类别，需要通过body 进行判断
	//继续修正，需要把全部数字去除
	subject = utils.DelDigitalInString(subject)
	partition, tag := ac.GetCategoryIdx(subject)
	if partition != utils.UnknownCategory {
		return partition, tag
	}
	newBody := e.AmendSubject(body)
	return ac.GetCategoryIdx(newBody)
}

//AuditEmailLegality ...基于解析内容判断邮件是否合法
func (e estimate) AuditEmailLegality(eml *model.Body, amendSubject, subjectTag string) utils.LegalTag {
	//步骤1：通过发件者的邮件域名，如果白名单直接为合法
	senderDomain := utils.GetSenderDomain(eml.From)
	v, ok := e.domainWhite[senderDomain]
	if ok {
		return v
	}
	//步骤2：通过标题中识别关键字，如果subjectTag不为空，判断通过关键字是否可以确定其为异常
	if subjectTag != "" {
		val, ok := utils.TagProperty[subjectTag]
		if ok && val == utils.InvalidTag {
			return utils.InvalidTag
		}
	}

	content := eml.Subject + eml.Body
	amendContent := e.AmendBody(content)
	//第3部：判断body中是否包括白名单数据，如jd.com 或者官方客服电话800-
	whiteWords := ac.GetWhiteHighlights(content) //使用原始数据，不要做修正
	if len(whiteWords) > 0 {
		return utils.ValidTag
	}

	//步骤4：提取微信号和QQ号,识别前需要做内容做修正，如提出空格，各类括号等内容。
	vxIDs := utils.GetVX(amendContent)
	qqIDs := utils.GetQQ(amendContent)
	fmt.Println(vxIDs, qqIDs)
	if len(vxIDs) > 0 || len(qqIDs) > 0 {
		fmt.Println("---------by 微信 qq-----------------")
		return utils.InvalidTag
	}
	//第5步骤：body直接已某些关键字为开头的，直接判断为合法
	if strings.HasPrefix(eml.Body, "发自我的") || strings.HasPrefix(eml.Body, "sentfrommy") {
		return utils.ValidTag
	}
	//第5步骤：判断正文中是否包括知名企业的域名，如JD.com,cebbank.com，就判断为合法

	//第6步骤：如果标题中出现手机号，并且类别是发票类，直接判断为非法邮件
	aaa := e.AmendBody(eml.Subject)
	mobilePhoneIDs, _ := utils.ExtractMobilePhone(aaa)
	if len(mobilePhoneIDs) > 0 {
		fmt.Println("----标题中有手机号--------------")
		return utils.InvalidTag
	}

	return utils.ValidTag
}

// AuditAllEmailItems  ...获取待鉴别邮件的分类
func (e estimate) AuditAllEmailItems() error {
	items, err := model.BodyModel.GetAllItems()
	if err != nil {
		return nil
	}
	for _, v := range items {
		v.Body = strings.ToLower(v.Body)
		v.Subject = strings.ToLower(v.Subject)
		amendSubject := e.AmendSubjectForCategory(v.Subject)
		//先计算其分类，然后更新到数据库中，后续可以比较了存入数据的分类是否和计算的分类一致。
		partition, tag := e.GetCategory(amendSubject, v.Body)
		v.Partition = partition.Name()
		err = model.BodyModel.UpdateItemCols(v, []string{"partition"})
		if err != nil {
			return err
		}
		//计算邮件是否异常
		v.ValidCalculate = int(e.AuditEmailLegality(v, amendSubject, tag))
		err = model.BodyModel.UpdateItemCols(v, []string{"valid_calculate"})
		if err != nil {
			return err
		}

	}
	return nil
}
