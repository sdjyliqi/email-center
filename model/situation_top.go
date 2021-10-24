package model

import (
	"email-center/utils"
	"github.com/golang/glog"
	"time"
)

var SituationTopModel SituationTop

type SituationTop struct {
	Id           int       `json:"id" xorm:"not null pk INT(11)"`
	Category     string    `json:"category" xorm:"VARCHAR(255)"`
	Name         string    `json:"name" xorm:"not null comment('分类的名称') VARCHAR(128)"`
	Amount       int       `json:"amount" xorm:"comment('数量') INT(11)"`
	LastModified time.Time `json:"last_modified" xorm:"DATE"`
}

func (t SituationTop) TableName() string {
	return "situation_top"
}

//GetAllItems ... 按照分类获取数据，
func (t SituationTop) GetAllItems(category string) ([]*SituationTop, error) {
	var items []*SituationTop
	err := utils.GetMysqlClient().Where("category=?", category).Desc("amount").Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
