package model

import (
	"email-center/utils"
	"github.com/golang/glog"
	"time"
)

var SituationPeriodModel SituationPeriod

type SituationPeriod struct {
	Id           int       `json:"id" xorm:"not null pk INT(11)"`
	Event        time.Time `json:"event" xorm:"comment('时间发送的日期') DATE"`
	Category     string    `json:"category" xorm:"VARCHAR(255)"`
	Name         string    `json:"name" xorm:"not null comment('分类的名称') VARCHAR(16)"`
	Amount       int       `json:"amount" xorm:"comment('数量') INT(11)"`
	LastModified time.Time `json:"last_modified" xorm:"DATE"`
}

func (t SituationPeriod) TableName() string {
	return "situation_period"
}

//GetAllItems ...获取N天的数据。
func (t SituationPeriod) GetAllItems(days int) ([]*SituationPeriod, error) {
	var items []*SituationPeriod
	err := utils.GetMysqlClient().OrderBy("event").Where("event >?", utils.GetDaysAgo(days)).Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
