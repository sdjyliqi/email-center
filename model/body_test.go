package model

import "testing"

func TestBody_GetAllItemsGetAllItems(t *testing.T) {
	items, err := BodyModel.GetAllItems()
	t.Log(items, err)
}
