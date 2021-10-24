package utils

import (
	"encoding/json"
	"sort"
	"testing"
)

func TestStringSlice(t *testing.T) {
	s := []string{"a", "你", "abc", "ab", "我的名字", "123455"}
	sort.Sort(StringSlice(s))
	t.Log(s)

}

func TestDelTo(t *testing.T) {
	var s []string
	contents := "[\"gaorongjun@wh.cebbank.com\"]"
	err := json.Unmarshal([]byte(contents), &s)
	t.Log(err, s)

	s2 := []string{"aaaa", "bbbbb"}

	c, err := json.Marshal(s2)
	t.Log(string(c), err)

}
