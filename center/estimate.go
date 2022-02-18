package center

import (
	"email-center/ac"
	"email-center/model"
	"email-center/utils"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"strings"
	"time"
)

//Estimate ...判断维度是否是否为真假
type Estimate struct {
	assistCharacter []string //设置
	amendCharacters []*model.Amend
	domainBillWhite map[string]utils.LegalTag
	domainADWhite   map[string]utils.LegalTag
	amendDict       map[string]string
}

//CreateEstimate ... 创建一个鉴定实例
func CreateEstimate() (*Estimate, error) {
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

	//初始化一些AC自动机的内容
	var replaceItems []string
	amendDict := map[string]string{}
	for _, v := range amendItems {
		replaceItems = append(replaceItems, v.Raw)
		amendDict[v.Raw] = v.Replace
	}
	ac.InitReplaceCharAC(replaceItems)
	return &Estimate{
		assistCharacter: characters,
		amendCharacters: amendItems,
		domainBillWhite: domainsForBill,
		domainADWhite:   domainsForAD,
		amendDict:       amendDict,
	}, nil
}

func (e Estimate) ReplaceContentByAC(content string) string {
	hitWords := ac.GetReplaceCharWords(content)
	for _, v := range hitWords {
		val, _ := e.amendDict[v]
		content = strings.ReplaceAll(content, v, val)
	}
	return content
}

//AmendSubject ...修正标题,剔除一些无用的字符
func (e Estimate) AmendSubjectForCategory(content string) string {
	var amendChars []rune
	chars := []rune(content)
	for _, v := range chars {
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
func (e Estimate) AmendSubject(content string) string {
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
func (e Estimate) AmendSubjectExtent(content string) string {
	newSubject := content
	for _, v := range e.amendCharacters {
		newSubject = strings.ReplaceAll(newSubject, v.Raw, v.Replace)
	}
	return newSubject
}

func (e Estimate) AmendRemoveReceive(eml *model.Body, content string) (string, error) {
	var receiver []string
	eml.To = strings.ReplaceAll(eml.To, "'", "\"")
	if len(eml.To) > 0 {
		err := json.Unmarshal([]byte(eml.To), &receiver)
		if err != nil {
			glog.Errorf("Unmarshal the %s failed,err:%+v", eml.To, err)
			return content, err
		}
	}
	for _, v := range receiver {
		content = strings.ReplaceAll(content, v, "")
	}
	return content, nil
}

//AmendBody ...修正邮件正文
func (e Estimate) AmendBody(content string) string {
	var amendChars []rune
	chars := []rune(content)
	for _, v := range chars {
		if v < '0' || v > '9' && v < 'A' || v > 'Z' && v < 'a' || v > 'z' && v <= 255 {
			continue
		}
		amendChars = append(amendChars, v)
	}
	newContent := string(amendChars)
	newContent = e.ReplaceContentByAC(newContent)

	for _, v := range e.assistCharacter {
		newContent = strings.ReplaceAll(newContent, v, "")
	}
	newContent = e.ReplaceContentByAC(newContent)
	//修补一下如果是com的情况，可能已经被替换为c0m，需要重新替换一下
	if strings.IndexAny(newContent, "c0m") > 0 {
		newContent = strings.ReplaceAll(newContent, "c0m", "com")
	}
	return newContent
}

//GetCategory ...获取待鉴别邮件的分类,标题、附件名、内容
func (e Estimate) GetCategory(subject, attachments, body string) (utils.Category, string) {
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
func (e Estimate) AuditEmailLegality(eml *model.Body, amendSubject, subjectTag string, category utils.Category) utils.LegalTag {
	switch category {
	case utils.BillCategory:
		return e.AuditBillEmail(eml, amendSubject, subjectTag)
	case utils.AdvertCategory:
		return e.AuditAdvEmail(eml, amendSubject, subjectTag)
	case utils.DirtyCategory:
		return e.AuditDirtyEmail(eml, amendSubject, subjectTag)
	default:
		return utils.UnknownTag
	}
	return utils.UnknownTag
}

//AuditBillEmail ...基于解析内容判断邮件是否合法
func (e Estimate) AuditBillEmail(eml *model.Body, amendSubject, subjectTag string) utils.LegalTag {
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
	amendContent := amendSubject + e.AmendBody(eml.Body)
	amendContent, _ = e.AmendRemoveReceive(eml, amendContent)
	//第3部：判断body中是否包括白名单数据，如jd.com 或者官方客服电话800-
	whiteWords := ac.GetWhiteHighlights(amendContent) //使用原始数据，不要做修正
	if len(whiteWords) > 0 {
		return utils.ValidTag
	}
	//步骤4：提取微信号和QQ号,识别前需要做内容做修正，如提出空格，各类括号等内容。
	vxIDs := utils.GetVX(amendSubject + amendContent)
	qqIDs := utils.GetQQ(amendSubject + amendContent)
	phoneIDs, _ := utils.ExtractMobilePhone(amendSubject + amendContent)
	fmt.Println("====", vxIDs, qqIDs, phoneIDs)
	if len(vxIDs) > 0 || len(qqIDs) > 0 || len(phoneIDs) > 0 {
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
		return utils.InvalidTag
	}
	//第7步：举个例子，标题中出现“发票”，正文中出现“代开发票”，判断类别时，标题中命中“发票”，直接判为发票类，内容不检查，后续判断合、非法时又没有检索关键字“代开发票”的步骤
	//导致很明显的代开发票邮件误判，故最后一步再检查一下关键词，
	amendBody := e.AmendBody(eml.Body)
	partition, tag := ac.GetCategoryIdx(amendBody)
	if partition == utils.BillCategory && utils.TagBillProperty[tag] == utils.InvalidTag {
		return utils.InvalidTag
	}
	return utils.UnknownTag
}

//AuditAdvEmail ...判断广告类邮件是否合法
func (e Estimate) AuditAdvEmail(b *model.Body, amendSubject, subjectTag string) utils.LegalTag {
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
	//第4步骤：body直接已某些关键字为开头的，直接判断为合法
	if strings.HasPrefix(b.Body, "发自我的") || strings.HasPrefix(b.Body, "sentfrommy") {
		return utils.ValidTag
	}
	//步骤5，在广告的分类下，出现微信等黑名单词，直接判断为异常状态
	blackWords := ac.GetADBlackWords(b.Body + amendSubject)
	if len(blackWords) > 0 {
		return utils.InvalidTag
	}
	return utils.UnknownTag
}

//AuditDirtyEmail ...判断色情类邮件
func (e Estimate) AuditDirtyEmail(b *model.Body, amendSubject, subjectTag string) utils.LegalTag {
	//步骤1：通过发件者的邮件域名，如果白名单直接为合法,2021年10月20日，临时取消该策略
	//senderDomain := utils.GetSenderDomain(b.From)
	//v, ok := e.domainADWhite[senderDomain]
	//if ok {
	//	return v
	//}
	//步骤2：判断色情灰词的数量和整体占比，
	amendContent := e.AmendBody(b.Body)
	dirtyWords := ac.GetDirtyWords(amendSubject + amendContent)
	if len(dirtyWords) >= 3 {
		return utils.InvalidTag
	}
	//判断是否包括：域名，qq，微信，tele
	domains, ok := utils.ExtractWebDomain(b.Body + amendSubject)
	if ok && len(domains) > 0 {
		return utils.InvalidTag
	}
	qqIDs := utils.GetQQ(amendSubject + amendContent)
	if len(qqIDs) > 0 {
		return utils.InvalidTag
	}
	vxIDs := utils.GetVX(amendSubject + amendContent)
	if len(vxIDs) > 0 {
		return utils.InvalidTag
	}
	return utils.UnknownTag
}

// AuditAllEmailItems  ...获取待鉴别邮件的分类
func (e Estimate) AuditAllEmailItems() error {
	fmt.Println("开始计算待审计的邮件内容")
	fmt.Println(time.Now())
	condition := " is_identify=0 "
	cnt, err := model.BodyModel.GetItemsCount(condition)
	if err != nil {
		return err
	}
	fmt.Println("待审计邮件的条数为：", cnt)
	items, err := model.BodyModel.GetItemsByCondition(condition)
	if err != nil {
		return err
	}
	e.AuditAllEmailItemsByPage(items)
	return nil
}

// AuditAllEmailItems  ...获取待鉴别邮件的分类
func (e Estimate) AuditAllEmailItemsByPage(items []*model.Body) error {
	for _, v := range items {
		fmt.Println("==========begin 处理：", v.Id)
		v.Body = strings.ToLower(v.Body)
		v.Subject = strings.ToLower(v.Subject)
		v.Attachments = strings.ToLower(v.Attachments)
		amendSubject := e.AmendSubjectForCategory(v.Subject)
		amendAttachments := e.AmendSubjectForCategory(v.Attachments)
		//先计算其分类，然后更新到数据库中，后续可以比较了存入数据的分类是否和计算的分类一致。
		partition, tag := e.GetCategory(amendSubject, amendAttachments, v.Body)
		v.Partition = partition.Name()
		err := model.BodyModel.UpdateItemCols(v, []string{"partition"})
		if err != nil {
			return err
		}
		//计算邮件是否异常
		v.ValidCalculate = int(e.AuditEmailLegality(v, amendSubject, tag, partition))
		v.IsIdentify = 1
		err = model.BodyModel.UpdateItemCols(v, []string{"valid_calculate", "is_identify"})
		if err != nil {
			return err
		}
	}
	fmt.Println(time.Now())
	return nil
}
