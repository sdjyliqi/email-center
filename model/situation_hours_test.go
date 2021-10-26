package model

import "testing"

func TestSituationHours_GetAllItems(t *testing.T) {
	items, err := SituationHoursModel.GetAllItems()
	t.Log(items, err)
	for _, v := range items {
		t.Log(v)
	}
}
