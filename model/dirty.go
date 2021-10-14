package model

import (
	"email-center/utils"
	"github.com/golang/glog"
)

//定义一个违禁词的实例
var DirtyModel Dirty

type Dirty struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Word     string `json:"word" xorm:"not null unique CHAR(128)"`
	Category string `json:"category" xorm:"default '' VARCHAR(128)"`
}

//设置表的名称
func (t Dirty) TableName() string {
	return "dirty"
}

//GetAllItems 获取全量的数据
func (t Dirty) GetAllItems() ([]*Dirty, error) {
	var items []*Dirty
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
