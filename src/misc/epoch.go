package misc

/*

References:
https://stackoverflow.com/questions/43915900/how-to-convert-unix-time-to-time-time-in-golang

*/

import (
	"time"
)

func DateToEpoch(year, month, day, hour, minute, second int) int {
	date := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
	return int(date.Unix())
}

func EpochToDate(epoch int) string {
	result := time.Unix(int64(epoch), 0)
	return result.Format("2006-1-2 15:4:5")
}
