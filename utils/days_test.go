package utils

import "testing"

func Test_GetDaysAgo(t *testing.T) {
	info := GetDaysAgo(30)
	t.Log(info)
}
