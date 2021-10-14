package model

import "testing"

func TestPortion_GetIdxItems(t *testing.T) {
	items, err := PortionModel.GetIdxItems("kai")
	t.Log(items, err)
	for _, v := range items {
		chars := []rune(v.Word)
		t.Log(v.Id, chars[0])
	}
}

func TestPortion_GetCategoryOfIdx(t *testing.T) {
	items, err := PortionModel.GetCategoryOfIdx()
	t.Log(items, err)
	for _, v := range items {
		chars := []rune(v.Idx)
		t.Log(string(chars))
	}
}
