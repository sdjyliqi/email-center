package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SituationPeriodModelGetAllItems(t *testing.T) {
	items, err := SituationPeriodModel.GetAllItems(30)
	assert.Nil(t, err)
	t.Log(items, err)
	for _, v := range items {
		t.Log(v)
	}
}
