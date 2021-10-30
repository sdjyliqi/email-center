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

//GetItemsByPage...  pageID页面从0开始
func (t AssistCharacter) GetItemsByPage(pageID, entry int) ([]*AssistCharacter, error) {
	var items []*AssistCharacter
	err := utils.GetMysqlClient().Limit(entry, pageID*entry).Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemsCount...  获取数量
func (t AssistCharacter) GetItemsCount() (int64, error) {
	cnt, err := utils.GetMysqlClient().Count(&AssistCharacter{})
	if err != nil {
		glog.Errorf("Get amount of items from %s failed,err:%+v", t.TableName(), err)
		return 0, err
	}
	return cnt, nil
}

//SearchItemsByIdx...  pageID页面从0开始
func (t AssistCharacter) SearchItemsByIdx(idx string) ([]*AssistCharacter, error) {
	var items []*AssistCharacter
	err := utils.GetMysqlClient().Where(" character like ?", "%"+idx+"%").Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//DelItemByID...  pageID页面从0开始
func (t AssistCharacter) DelItemByID(id int64) error {
	_, err := utils.GetMysqlClient().ID(id).Delete(AssistCharacter{})
	if err != nil {
		glog.Errorf("Delete item by id %d from %s failed,err:%+v", id, t.TableName(), err)
		return err
	}
	return nil
}

func (t AssistCharacter) UpdateItemByID(item *AssistCharacter) error {
	item.Lastmodified = time.Now()
	if item.Id > 0 {
		cols := []string{"character"}
		_, err := utils.GetMysqlClient().Cols(cols...).ID(item.Id).Update(item)
		if err != nil {
			glog.Errorf("Update item %+v from %s failed,err:%+v", *item, t.TableName(), err)
			return err
		}
		return nil
	}
	item.Id = 0
	_, err := utils.GetMysqlClient().Insert(item)
	if err != nil {
		glog.Errorf("insert item %+v into table %s failed,err:%+v", *item, t.TableName(), err)
		return err
	}
	return nil
}
