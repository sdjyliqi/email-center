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
func (e estimate) AmendSubject(content string) string {
	var amendChars []rune
	chars := []rune(content)
	for _, v := range chars {
		if (v < 'A') || (v > 'z' && v <= 255) || (v == '\\') {
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

//AmendBody ...修正邮件正文
func (e estimate) AmendBody(content string) string {
	var amendChars []rune
	chars := []rune(content)
	for _, v := range chars {
		if v < '0' || v > 'z' && v <= 255 {
			continue
		}
		amendChars = append(amendChars, v)
	}
	newSubject := string(amendChars)
	for _, v := range e.assistCharacter {
		newSubject = strings.ReplaceAll(newSubject, v, "")
	}
	for _, v := range e.amendCharacters {
		newSubject = strings.ReplaceAll(newSubject, v.Raw, v.Replace)
	}
	return newSubject
}

//GetCategory ...获取待鉴别邮件的分类
func (e estimate) GetCategory(subject, body string) (utils.Category, string) {
	newSubject := e.AmendSubject(subject)
	fmt.Println("================", subject, newSubject)
	//如果通过标题无法判断出类别，需要通过body 进行判断
	partition, tag := utils.GetCategoryIdx(newSubject)
	if partition != utils.UnknownCategory {
		return partition, tag
	}
	newBody := e.AmendSubject(body)
	return utils.GetCategoryIdx(newBody)
}

//AuditEmailLegality ...基于解析内容判断邮件是否合法
func (e estimate) AuditEmailLegality(eml *model.Body, subjectTag string) utils.LegalTag {
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
	//步骤3：提取微信号和QQ号,识别前需要做内容做修正，如提出空格，各类括号等内容。
	content := eml.Subject + eml.Body
	amendContent := e.AmendBody(content)
	vxIDs := utils.GetVX(amendContent)
	qqIDs := utils.GetQQ(amendContent)
	if len(vxIDs) > 0 || len(qqIDs) > 0 {
		return utils.InvalidTag
	}
	//第4步骤：body直接已某些关键字为开头的，直接判断为合法
	if strings.HasPrefix(eml.Body, "发自我的") || strings.HasPrefix(eml.Body, "sentfrommy") {
		return utils.ValidTag
	}
	//第5步骤：判断正文中是否包括知名企业的域名，如JD.com,cebbank.com，就判断为合法

	//第6步骤：如果标题中出现手机号，并且类别是发票类，直接判断为非法邮件
	if eml.Partition == utils.BillCategory.Name() {
		mobilePhoneIDs, _ := utils.ExtractMobilePhone(eml.Subject)
		if len(mobilePhoneIDs) > 1 {
			return utils.InvalidTag
		}
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
		//先计算其分类，然后更新到数据库中，后续可以比较了存入数据的分类是否和计算的分类一致。
		partition, tag := e.GetCategory(v.Subject, v.Body)
		v.Partition = partition.Name()
		err = model.BodyModel.UpdateItemCols(v, []string{"partition"})
		if err != nil {
			return err
		}
		//计算邮件是否异常
		v.ValidCalculate = e.AuditEmailLegality(v, tag)
		err = model.BodyModel.UpdateItemCols(v, []string{"valid_calculate"})
		if err != nil {
			return err
		}

	}
	return nil
}
