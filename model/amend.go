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

//GetItemsByPage...  pageID页面从0开始
func (t Amend) GetItemsByPage(pageID, entry int) ([]*Amend, error) {
	var items []*Amend
	err := utils.GetMysqlClient().Limit(entry, pageID*entry).Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetItemsCount...  获取数量
func (t Amend) GetItemsCount() (int64, error) {
	cnt, err := utils.GetMysqlClient().Count(&Dirty{})
	if err != nil {
		glog.Errorf("Get amount of items from %s failed,err:%+v", t.TableName(), err)
		return 0, err
	}
	return cnt, nil
}

//SearchItemsByIdx...  pageID页面从0开始
func (t Amend) SearchItemsByIdx(idx string) ([]*Amend, error) {
	var items []*Amend
	err := utils.GetMysqlClient().Where(" raw like ?", "%"+idx+"%").Find(&items)
	if err != nil {
		glog.Errorf("Get items from %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//DelItemByID...  pageID页面从0开始
func (t Amend) DelItemByID(id int64) error {
	_, err := utils.GetMysqlClient().ID(id).Delete(Amend{})
	if err != nil {
		glog.Errorf("Delete item by id %d from %s failed,err:%+v", id, t.TableName(), err)
		return err
	}
	return nil
}

func (t Amend) UpdateItemByID(item *Amend) error {
	item.Lastmodified = time.Now()
	if item.Id > 0 {
		cols := []string{"raw", "replace"}
		_, err := utils.GetMysqlClient().Cols(cols...).ID(item.Id).Update(item)
		if err != nil {
			glog.Errorf("Update item %+v from %s failed,err:%+v", *item, t.TableName(), err)
			return err
		}
		return nil
	}
	item.Id = 0
	//修改色情分类
	_, err := utils.GetMysqlClient().Insert(item)
	if err != nil {
		glog.Errorf("insert item %+v into table %s failed,err:%+v", *item, t.TableName(), err)
		return err
	}
	return nil
}
