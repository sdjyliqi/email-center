package model

import (
	"email-center/utils"
	"github.com/golang/glog"
	"time"
)

var AmendModel = Amend{}

type Amend struct {
	Id           int       `json:"id" xorm:"not null pk INT(11)"`
	Raw          string    `json:"raw" xorm:"not null unique VARCHAR(1)"`
	Replace      string    `json:"replace" xorm:"VARCHAR(1)"`
	Lastmodified time.Time `json:"lastmodified" xorm:"DATE"`
}

func (t Amend) TableName() string {
	return "amend"
}

//GetAllItems ...
func (t Amend) GetAllItems() ([]*Amend, error) {
	var items []*Amend
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
