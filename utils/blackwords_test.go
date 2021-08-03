package utils

import (
	"encoding/json"
	"testing"
)

func Test_CreateFAPIAOBlackIndex(t *testing.T) {
	result := CreateFAPIAOBlackIndex()
	t.Log(result)
	content, _ := json.Marshal(result)
	t.Log(string(content))
}
