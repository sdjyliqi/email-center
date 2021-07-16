package model

import (
	"email-center/utils"
	"github.com/golang/glog"
	"strings"
)

var BodyModel Body

type Body struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	FileName        string `json:"file_name" xorm:"unique VARCHAR(128)"`
	SendTime        string `json:"send_time" xorm:"VARCHAR(64)"`
	From            string `json:"from" xorm:"not null VARCHAR(128)"`
	To              string `json:"to" xorm:"not null VARCHAR(4096)"`
	Subject         string `json:"subject" xorm:"VARCHAR(256)"`
	Category        string `json:"category" xorm:"VARCHAR(32)"`
	ContentLanguage string `json:"content_language" xorm:"VARCHAR(16)"`
	MessageId       string `json:"message_id" xorm:"VARCHAR(128)"`
	ContentLength   int    `json:"content_length" xorm:"INT(11)"`
	Body            string `json:"body" xorm:"VARCHAR(10240)"`
}

func (t Body) TableName() string {
	return "body"
}

//NoExistedInsertItem ...插入数据库数据，前提是基于message id不存在
func (t *Body) NoExistedInsertItem(item *utils.Email) error {
	var emailItem = Body{}
	ok, err := utils.GetMysqlClient().Where("message_id=?", item.MessageID).Get(&emailItem)
	if err != nil {
		glog.Errorf("Get item by message_id %s from table %s failed,err:%+v", item.MessageID, t.TableName(), err)
		return err
	}
	item.ContentBody = strings.Replace(item.ContentBody, "\r", "", -1)
	item.ContentBody = strings.Replace(item.ContentBody, "\n", "", -1)
	if !ok {
		var newItem = Body{
			FileName:        item.FileName,
			From:            item.From,
			To:              item.To,
			Subject:         item.Subject,
			Category:        item.Category,
			ContentLanguage: "",
			MessageId:       item.MessageID,
			ContentLength:   0,
			Body:            item.ContentBody,
			SendTime:        item.Date,
		}
		_, err = utils.GetMysqlClient().Insert(newItem)
		if err != nil {
			glog.Errorf("Insert the to table %s failed,err:%+v", t.TableName(), err)
			return err
		}
	}
	return nil
}
