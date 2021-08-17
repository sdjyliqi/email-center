package model

import "testing"

func TestSensitive_GetAllItems(t *testing.T) {
	items, err := SensitiveModel.GetAllItems()
	t.Log(items, err)
	for _, v := range items {
		chars := []rune(v.Words)
		t.Log(v.Id, string(chars))
	}
}
