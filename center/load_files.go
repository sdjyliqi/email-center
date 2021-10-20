package center

import (
	"email-center/ac"
	"email-center/model"
	"email-center/utils"
	"fmt"
)

type EmailFile struct {
	path string         //设置邮件分类的目录
	name string         //设置邮件分类的名称
	tag  utils.LegalTag //按照目录导入的时候，可能已经对内容进行了合法性分类
}

type LoadEmailBody struct {
	emailItems []EmailFile
}

//CreateLoadEmailBody 创建结构体
func CreateLoadEmailBody() *LoadEmailBody {
	return &LoadEmailBody{
		emailItems: []EmailFile{},
	}
}

//SetDataPath ...设置数据文件配置信息
func (l *LoadEmailBody) SetDataPath(item []EmailFile) error {
	l.emailItems = item
	return nil
}

func (l *LoadEmailBody) ExtractEmailData() error {
	items, err := model.BodyModel.GetAllItems()
	fmt.Println("===数量：========", len(items))
	if err != nil {
		return err
	}
	for _, v := range items {
		body := model.Extract{Id: v.Id}
		sender := v.From
		senderDomain := utils.GetSenderDomain(sender)
		body.SenderDomain = senderDomain
		weixinIDs := utils.GetVX(v.Body)
		if len(weixinIDs) > 0 {
			body.Weixin = weixinIDs[0]
		}
		qqIDs := utils.GetQQ(v.Body)
		fmt.Println("weixin:", weixinIDs, "QQ:", qqIDs)
		if len(qqIDs) > 0 {
			body.Qq = qqIDs[0]
		}
		bodyURLDomains, err := utils.ExtractWebDomain(v.Body)
		fmt.Println("======", bodyURLDomains, err)

		//寻找类别关键字，
		categoryWords := ac.CategoryACMatch.Match(v.Subject + v.Body + v.Attachments)
		fmt.Println(categoryWords)
		//fmt.Println("=============",categoryWords)
	}
	return nil
}
