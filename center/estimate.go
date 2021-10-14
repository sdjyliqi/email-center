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
	domainBillWhite map[string]utils.LegalTag
	domainADWhite   map[string]utils.LegalTag
}

// CreateEstimate ... 创建一个鉴定实例
func CreateEstimate() (*estimate, error) {
	var characters []string
	domainsForBill := map[string]utils.LegalTag{}
	domainsForAD := map[string]utils.LegalTag{}
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
		//todo 默认全部域名为发票类的白名单
		domainsForBill[v.Official] = utils.ValidTag
		if v.AllowAd == 1 {
			domainsForAD[v.Official] = utils.ValidTag
		}
	}
	//
	amendItems, err := model.AmendModel.GetAllItems()
	if err != nil {
		return nil, err
	}

	return &estimate{
		assistCharacter: characters,
		amendCharacters: amendItems,
		domainBillWhite: domainsForBill,
		domainADWhite:   domainsForAD,
	}, nil
}

//AmendSubject ...修正标题,剔除一些无用的字符
func (e estimate) AmendSubjectForCategory(content string) string {
	var amendChars []rune
	chars := []rune(content)
	for _, v := range chars {
		//v < 'A' || v > 'Z' && v < 'a' || v > 'z' && v <= 255 此处修改前的代码
		if v < '0' || v > '9' && v < 'A' || v > 'Z' && v < 'a' || v > 'z' && v <= 255 {
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

//GetCategory ...获取待鉴别邮件的分类,标题、附件名、内容
func (e estimate) GetCategory(subject, attachments, body string) (utils.Category, string) {
	//如果通过标题无法判断出类别，需要通过body 进行判断
	//继续修正，需要把全部数字去除
	subjectNoNum := utils.DelDigitalInString(subject)
	attachmentsNoNum := utils.DelDigitalInString(attachments)
	partition, tag := ac.GetCategoryIdx(subjectNoNum + attachmentsNoNum)

	if partition != utils.UnknownCategory {
		return partition, tag
	}

	//因为5折，这样的数字会被去掉，导致广告类关键词遭到破坏，又因为发票类关键词比广告类准确，所以只能先去除数字判断bill类，再保留数字判断ad类
	partition, tag = ac.GetCategoryIdx(subject + attachments)
	if partition != utils.UnknownCategory {
		return partition, tag
	}
	newBody := e.AmendSubject(body)
	return ac.GetCategoryIdx(newBody)
}

//AuditEmailLegality ...基于解析内容判断邮件是否合法
//传入参数 eml为邮件内容结构图，amendSubject修正后的主题，通过分类的时候命中的分类标签
func (e estimate) AuditEmailLegality(eml *model.Body, amendSubject, subjectTag string, category utils.Category) utils.LegalTag {
	switch category {
	case utils.BillCategory:
		return e.AuditBillEmail(eml, amendSubject, subjectTag)
	case utils.AdvertCategory:
		return e.AuditAdvEmail(eml, amendSubject, subjectTag)
	default:
		return utils.UnknownTag
	}
	return utils.UnknownTag
}

//AuditBillEmail ...基于解析内容判断邮件是否合法
func (e estimate) AuditBillEmail(eml *model.Body, amendSubject, subjectTag string) utils.LegalTag {
	//步骤1：通过发件者的邮件域名，如果白名单直接为合法
	senderDomain := utils.GetSenderDomain(eml.From)
	v, ok := e.domainBillWhite[senderDomain]
	if ok {
		return v
	}
	//步骤2：通过标题中识别关键字，如果subjectTag不为空，判断通过关键字是否可以确定其为异常
	if subjectTag != "" {
		val, ok := utils.TagBillProperty[subjectTag]
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
	//第7步：举个例子，标题中出现“发票”，正文中出现“代开发票”，判断类别时，标题中命中“发票”，直接判为发票类，内容不检查，后续判断合、非法时又没有检索关键字“代开发票”的步骤
	//导致很明显的代开发票邮件误判，故最后一步再检查一下关键词，
	amendbody := e.AmendBody(eml.Body)
	partition, tag := ac.GetCategoryIdx(amendbody)
	if partition == utils.BillCategory && utils.TagBillProperty[tag] == utils.InvalidTag {
		return utils.InvalidTag
	}
	return utils.ValidTag
}

//AuditAdvEmail ...判断广告类邮件是否合法
func (e estimate) AuditAdvEmail(b *model.Body, amendSubject, subjectTag string) utils.LegalTag {
	//步骤1：通过发件者的邮件域名，如果白名单直接为合法
	senderDomain := utils.GetSenderDomain(b.From)
	v, ok := e.domainADWhite[senderDomain]
	if ok {
		return v
	}
	//步骤2：通过标题中识别关键字，如果subjectTag不为空，判断通过关键字是否可以确定其为异常
	val, ok := utils.TagADProperty[subjectTag]
	if ok && val == utils.InvalidTag {
		return utils.InvalidTag
	}
	//步骤3：判断是否包括已入库单位的电话等信息，如果有直接为
	b.Body = e.AmendBody(b.Body)
	customerServiceIDs := ac.GetCustomerServiceIDs(b.Body + amendSubject)
	if len(customerServiceIDs) > 0 {
		return utils.ValidTag
	}
	//步骤4，在广告的分类下，出现微信等黑名单词，直接判断为异常状态
	blackWords := ac.GetADBlackWords(b.Body + amendSubject)
	if len(blackWords) > 0 {
		return utils.InvalidTag
	}
	return utils.UnknownTag
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
		v.Attachments = strings.ToLower(v.Attachments)
		amendSubject := e.AmendSubjectForCategory(v.Subject)
		amendattachments := e.AmendSubjectForCategory(v.Attachments)
		//先计算其分类，然后更新到数据库中，后续可以比较了存入数据的分类是否和计算的分类一致。
		partition, tag := e.GetCategory(amendSubject, amendattachments, v.Body)
		v.Partition = partition.Name()
		err = model.BodyModel.UpdateItemCols(v, []string{"partition"})
		if err != nil {
			return err
		}
		//计算邮件是否异常
		v.ValidCalculate = int(e.AuditEmailLegality(v, amendSubject, tag, partition))
		err = model.BodyModel.UpdateItemCols(v, []string{"valid_calculate"})
		if err != nil {
			return err
		}

	}
	return nil
}
