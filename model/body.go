package model

import (
	"email-center/utils"
	"encoding/json"
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
		glog.Errorf("Get item by filename %s from table %s failed,err:%+v", item.FileName, t.TableName(), err)
		return err
	}
	ok = false
	if !ok {
		fromContent, _ := json.Marshal(item.ParseEmail.From)
		toContent, _ := json.Marshal(item.ParseEmail.To)
		sendTime, _ := item.ParseEmail.Header.Date()
		strSendTime := sendTime.Format(utils.TimeFormat)
		body := item.Body
		var newItem = Body{
			FileName:        item.FileName,
			From:            string(fromContent),
			To:              string(toContent),
			Subject:         item.ParseEmail.Subject,
			Category:        item.Category,
			ContentLanguage: "",
			MessageId:       item.ParseEmail.MessageID,
			ContentLength:   len(body),
			Body:            body,
			SendTime:        strSendTime,
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
