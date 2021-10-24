package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SituationTopGetAllItems(t *testing.T) {
	items, err := SituationTopModel.GetAllItems("")
	assert.Nil(t, err)
	t.Log(items, err)
	for _, v := range items {
		t.Log(v)
	}
}
