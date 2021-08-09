package model

import (
	"email-center/utils"
	"github.com/golang/glog"
)

var BodyModel Body

type Body struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	FileName        string `json:"file_name" xorm:"unique VARCHAR(128)"`
	SendTime        string `json:"send_time" xorm:"VARCHAR(64)"`
	From            string `json:"from" xorm:"not null VARCHAR(128)"`
	To              string `json:"to" xorm:"not null VARCHAR(4096)"`
	Valid           int    `json:"valid" xorm:"TINYINT(4)"`
	Subject         string `json:"subject" xorm:"VARCHAR(256)"`
	Category        string `json:"category" xorm:"VARCHAR(32)"`
	ContentLanguage string `json:"content_language" xorm:"VARCHAR(16)"`
	MessageId       string `json:"message_id" xorm:"VARCHAR(128)"`
	ContentLength   int    `json:"content_length" xorm:"INT(11)"`
	Attachments     string `json:"attachments" xorm:"VARCHAR(1024)"`
	Body            string `json:"body" xorm:"TEXT"`
}

func (t Body) TableName() string {
	return "body_liqi"
}

//GetAllItems ...
func (t Body) GetAllItems() ([]*Body, error) {
	var items []*Body

	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
