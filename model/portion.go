package model

import (
	"email-center/utils"
	"fmt"
	"github.com/golang/glog"
	"time"
)

var PortionModel Portion

type Portion struct {
	Id           int       `json:"id" xorm:"not null pk INT(8)"`
	Word         string    `json:"word" xorm:"unique VARCHAR(8)"`
	Idx          string    `json:"idx" xorm:"VARCHAR(8)"`
	Category     string    `json:"category" xorm:"VARCHAR(8)"`
	Lastmodified time.Time `json:"lastmodified" xorm:"DATE"`
}

func (t Portion) TableName() string {
	return "portion"
}

func (t Portion) GetIdxItems(idx string) ([]*Portion, error) {
	var items []*Portion // Where("`category` != `partition` ")..Where("id=151").
	queryString := fmt.Sprintf("idx='%s'", idx)
	err := utils.GetMysqlClient().Where(queryString).Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

func (t Portion) GetCategoryOfIdx() ([]*Portion, error) {
	var items []*Portion
	err := utils.GetMysqlClient().SQL("SELECT idx FROM portion GROUP BY idx").Find(&items)
	if err != nil {
		glog.Errorf("Get CategoryOfIdx from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
