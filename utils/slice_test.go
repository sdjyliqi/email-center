package utils

import (
	"sort"
	"testing"
)

func TestStringSlice(t *testing.T) {
	s := []string{"a", "你", "abc", "ab", "我的名字", "123455"}
	sort.Sort(StringSlice(s))
	t.Log(s)

}
