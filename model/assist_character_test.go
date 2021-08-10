package model

import "testing"

func TestBody_AssistCharacterGetAllItems(t *testing.T) {
	items, err := AssistCharacterModule.GetAllItems()
	t.Log(items, err)
	for _, v := range items {
		chars := []rune(v.Character)
		if chars[0] < 255 {
			t.Log(v.Id, chars[0])
		}
	}
}
