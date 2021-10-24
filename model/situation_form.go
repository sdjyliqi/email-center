package model

import (
	"email-center/utils"
	"github.com/golang/glog"
	"time"
)

var SituationFormModel SituationForm

type SituationForm struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name         string    `json:"name" xorm:"not null comment('分类的名称') VARCHAR(16)"`
	Amount       int       `json:"amount" xorm:"comment('数量') INT(11)"`
	LastModified time.Time `json:"last_modified" xorm:"DATE"`
}

func (t SituationForm) TableName() string {
	return "situation_form"
}

//GetAllItems ...获取统计异常数据的比例图
func (t SituationForm) GetAllItems() ([]*SituationForm, error) {
	var items []*SituationForm
	err := utils.GetMysqlClient().Desc("amount").Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
