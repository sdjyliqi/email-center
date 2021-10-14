package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirty_GetAllItemsGetAllItems(t *testing.T) {
	items, err := DirtyModel.GetAllItems()
	assert.Nil(t, err)
	t.Log(items, err)
}
