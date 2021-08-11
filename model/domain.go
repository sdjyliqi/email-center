package model

import (
	"email-center/utils"
	"github.com/golang/glog"
)

var DomainModel = Domain{}

type Domain struct {
	Id       int    `json:"id" xorm:"not null pk autoincr comment('域名白名单自增id') INT(11)"`
	Suffix   string `json:"suffix" xorm:"not null comment('发件人邮箱后缀') unique VARCHAR(32)"`
	Name     string `json:"name" xorm:"not null comment('发件人邮箱简称') VARCHAR(32)"`
	Official string `json:"official" xorm:"comment('发件人公司名') VARCHAR(16)"`
}

func (t Domain) TableName() string {
	return "domain"
}

// GetAllItems 获取全量的数据，加工全面域名白名单
func (t Domain) GetAllItems() ([]*Domain, error) {
	var items []*Domain
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
