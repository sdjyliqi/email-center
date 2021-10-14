package model

import (
	"email-center/utils"
	"github.com/golang/glog"
	"time"
)

var SensitiveModel Sensitive

type Sensitive struct {
	Id           int       `json:"id" xorm:"not null pk INT(8)"`
	Words        string    `json:"words" xorm:"not null default '' unique VARCHAR(32)"`
	Lastmodified time.Time `json:"lastmodified" xorm:"DATE"`
}

func (t Sensitive) TableName() string {
	return "sensitive"
}

func (t Sensitive) GetAllItems() ([]*Sensitive, error) {
	var items []*Sensitive // Where("`category` != `partition` ")..Where("id=151").
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
