package model

import (
	"email-center/utils"
	"github.com/golang/glog"
)

//定义一个违禁词的实例
var DirtyModel Dirty

type Dirty struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Word         string    `json:"word" xorm:"not null unique CHAR(128)"`
	Category     string    `json:"category" xorm:"default '' VARCHAR(128)"`
	LastModified time.Time `json:"last_modified" xorm:"DATETIME"`
	Submit       string    `json:"submit" xorm:"CHAR(64)"`
}

//设置表的名称
func (t Dirty) TableName() string {
	return "dirty"
}

//GetAllItems 获取全量的数据
func (t Dirty) GetAllItems() ([]*Dirty, error) {
	var items []*Dirty
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemsByPage...  pageID页面从0开始
func (t Dirty) GetItemsByPage(pageID, entry int) ([]*Dirty, error) {
	var items []*Dirty
	err := utils.GetMysqlClient().Limit(entry, pageID*entry).Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemsCount...  获取数量
func (t Dirty) GetItemsCount() (int64, error) {
	cnt, err := utils.GetMysqlClient().Count(&Dirty{})
	if err != nil {
		glog.Errorf("Get amount of items from %s failed,err:%+v", t.TableName(), err)
		return 0, err
	}
	return cnt, nil
}

//SearchItemsByIdx...  pageID页面从0开始
func (t Dirty) SearchItemsByIdx(idx string) ([]*Dirty, error) {
	var items []*Dirty
	err := utils.GetMysqlClient().Where("category='色情' and word like ?", "%"+idx+"%").Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//DelItemByID...  pageID页面从0开始
func (t Dirty) DelItemByID(id int64) error {
	_, err := utils.GetMysqlClient().ID(id).Delete(Dirty{})
	if err != nil {
		glog.Errorf("Delete item by id %d from %s failed,err:%+v", id, t.TableName(), err)
		return err
	}
	return nil
}

func (t Dirty) UpdateItemByID(item *Dirty) error {
	if item.Id > 0 {
		cols := []string{"word"}
		_, err := utils.GetMysqlClient().Cols(cols...).ID(item.Id).Update(item)
		if err != nil {
			glog.Errorf("Update item %+v from %s failed,err:%+v", *item, t.TableName(), err)
			return err
		}
		return nil
	}
	item.Id = 0
	item.Category = "色情"
	//修改色情分类
	_, err := utils.GetMysqlClient().Insert(item)
	if err != nil {
		glog.Errorf("insert item %+v into table %s failed,err:%+v", *item, t.TableName(), err)
		return err
	}
	return nil
}
