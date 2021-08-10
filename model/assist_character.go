package model

import (
	"email-center/utils"
	"github.com/golang/glog"
	"time"
)

var AssistCharacterModule = &AssistCharacter{}

type AssistCharacter struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Character    string    `json:"character" xorm:"unique VARCHAR(16)"`
	Lastmodified time.Time `json:"lastmodified" xorm:"DATE"`
}

//TableName .....。设定与mysql的表名的关联
func (t AssistCharacter) TableName() string {
	return "assist_character"
}

// GetAllItems 获取全量的辅助字符，加工初始串
func (t AssistCharacter) GetAllItems() ([]*AssistCharacter, error) {
	var items []*AssistCharacter

	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
