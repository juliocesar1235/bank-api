package utilities

import (
	"time"
)

func GetMonthMap() map[string]time.Month {
	return map[string]time.Month{
		"Jan":  time.January,
		"Feb":  time.February,
		"Mar":  time.March,
		"Apr":  time.April,
		"May":  time.May,
		"Jun":  time.June,
		"Jul":  time.July,
		"Aug":  time.August,
		"Sep": 	time.September,
		"Oct":  time.October,
		"Nov":  time.November,
		"Dec":  time.December,
	}
}