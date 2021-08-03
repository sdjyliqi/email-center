package model

import (
	"email-center/utils"
	"github.com/golang/glog"
)

type Extract struct {
	Id           int    `json:"id" xorm:"not null pk INT(11)"`
	SenderDomain string `json:"sender_domain" xorm:"VARCHAR(64)"`
	BodyDomains  string `json:"body_domains" xorm:"default '' comment('识别出来的域名列表，中间用','分割') VARCHAR(1024)"`
	Keywords     string `json:"keywords" xorm:"comment('发票类关键字信息') VARCHAR(256)"`
	Phone        string `json:"phone" xorm:"VARCHAR(256)"`
	Weixin       string `json:"weixin" xorm:"VARCHAR(256)"`
	Qq           string `json:"qq" xorm:"VARCHAR(256)"`
}

func (t Extract) TableName() string {
	return "extract"
}

//GetAllItems ...
func (t Extract) GetAllItems() ([]*Extract, error) {
	var items []*Extract

	err := utils.GetMysqlClient().Find(items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
