package model

import (
	"email-center/utils"
	"errors"
	"github.com/golang/glog"
)

var BodyModel Body

type Body struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	FileName        string `json:"file_name" xorm:"unique VARCHAR(128)"`
	From            string `json:"from" xorm:"not null VARCHAR(128)"`
	IsIdentify      int    `json:"is_identify" xorm:"TINYINT(4)"`
	ValidManual     int    `json:"valid_manual" xorm:"comment('录入数据时候打的标记') TINYINT(4)"`
	ValidCalculate  int    `json:"valid_calculate" xorm:"TINYINT(4)"`
	SendTime        string `json:"send_time" xorm:"VARCHAR(64)"`
	To              string `json:"to" xorm:"not null VARCHAR(1024)"`
	Subject         string `json:"subject" xorm:"VARCHAR(256)"`
	Partition       string `json:"partition" xorm:"VARCHAR(32)"`
	Category        string `json:"category" xorm:"VARCHAR(32)"`
	ContentEncode   string `json:"content_encode" xorm:"VARCHAR(32)"`
	ContentLanguage string `json:"content_language" xorm:"VARCHAR(16)"`
	MessageId       string `json:"message_id" xorm:"VARCHAR(128)"`
	ContentLength   int    `json:"content_length" xorm:"INT(11)"`
	Attachments     string `json:"attachments" xorm:"VARCHAR(1024)"`
	Body            string `json:"body" xorm:"TEXT"`
}

func (t Body) TableName() string {
	return "body"
}

//GetAllItems ...i
func (t Body) GetAllItems() ([]*Body, error) {
	var items []*Body //.Where("is_identify=0").Where("is_identify=0"). Where("id=37")
	err := utils.GetMysqlClient().Where("valid_calculate != valid_manual").Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//UpdateItemCols ...根据ID 修改某列属性
func (t Body) UpdateItemCols(item *Body, cols []string) error {
	if item == nil {
		return errors.New("not-nil")
	}
	_, err := utils.GetMysqlClient().Cols(cols...).ID(item.Id).Update(item)
	if err != nil {
		glog.Errorf("Update the item %+v from %s failed,err:%+v", *item, t.TableName(), err)
		return err
	}
	return nil
}
