package model

import (
	"email-center/utils"
	"github.com/golang/glog"
	"time"
)

var SituationHoursModel SituationHours

type SituationHours struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name         string    `json:"name" xorm:"not null comment('分类的名称') VARCHAR(16)"`
	Amount       int       `json:"amount" xorm:"comment('数量') INT(11)"`
	LastModified time.Time `json:"last_modified" xorm:"DATE"`
}

func (t SituationHours) TableName() string {
	return "situation_hours"
}

//GetAllItems 获取全量的数据
func (t SituationHours) GetAllItems() ([]*SituationHours, error) {
	var items []*SituationHours
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
