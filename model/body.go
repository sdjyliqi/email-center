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
	Valid           int    `json:"valid" xorm:"TINYINT(4)"`
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

//NoExistedInsertItem ...插入数据库数据，前提是基于邮件路径不存在
func (t *Body) NoExistedInsertItem(item *utils.Email) error {
	var emailItem = Body{}
	ok, err := utils.GetMysqlClient().Where("file_name=?", item.FileName).Get(&emailItem)
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
			ContentLength:   len(item.ContentBody),
			Body:            item.ContentBody,
			SendTime:        item.Date,
			Valid:           int(item.Valid),
		}
		_, err = utils.GetMysqlClient().Insert(newItem)
		if err != nil {
			glog.Errorf("Insert the to table %s failed,err:%+v", t.TableName(), err)
			return err
		}
	}
	return nil
}
