package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDirty_GetAllItems(t *testing.T) {
	items, err := DirtyModel.GetAllItems()
	assert.Nil(t, err)
	for _, v := range items {
		t.Log(v)
	}
}

func TestDirty_GetItemsByPage(t *testing.T) {
	items, err := DirtyModel.GetItemsByPage(0, 10)
	assert.Nil(t, err)
	for _, v := range items {
		t.Log(v)
	}
}

func TestDirty_GetItemsCount(t *testing.T) {
	cnt, err := DirtyModel.GetItemsCount()
	assert.Nil(t, err)
	t.Log(cnt)
}

func TestDirty_SearchItemsByIdx(t *testing.T) {
	items, err := DirtyModel.SearchItemsByIdx("爱")
	assert.Nil(t, err)
	for _, v := range items {
		t.Log(v)
	}
}

func TestDirty_DelItemByID(t *testing.T) {
	err := DirtyModel.DelItemByID(time.Now().Unix())
	assert.Nil(t, err)

}
