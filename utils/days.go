package utils

import "time"

var DayCommonFormat = "20060102"

//GetDaysAgo ... 获取K天前的日期，days 为正数
func GetDaysAgo(days int) time.Time {
	t := time.Now()
	return t.AddDate(0, 0, -days)
}
