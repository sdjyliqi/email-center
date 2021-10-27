package model

import (
	"email-center/utils"
	"github.com/golang/glog"
	"time"
)

var SituationAmountModel SituationAmount

type SituationAmount struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Event        time.Time `json:"event" xorm:"DATE"`
	Abnormal     int       `json:"abnormal" xorm:"INT(11)"`
	Amount       int       `json:"amount" xorm:"comment('数量') INT(11)"`
	LastModified time.Time `json:"last_modified" xorm:"DATE"`
}

func (t SituationAmount) TableName() string {
	return "situation_amount"
}

//GetAllItems ...获取统计异常数据的比例图
func (t SituationAmount) GetAllItems() ([]*SituationAmount, error) {
	var items []*SituationAmount
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
