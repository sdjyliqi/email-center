package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SituationFormGetAllItems(t *testing.T) {
	items, err := SituationFormModel.GetAllItems()
	assert.Nil(t, err)
	t.Log(items, err)
	for _, v := range items {
		t.Log(v)
	}
}
