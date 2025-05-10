package util

// import "time"

// // DateISOFormat ...
// const DateISOFormat = "2006-01-02T15:04:05.000Z"

// const timezoneHCM = "Asia/Ho_Chi_Minh"

// // FormatTimeExcel ...
// const FormatTimeExcel = "2006-01-02 15:04:05"

// // TimeParseISODate ...
// func TimeParseISODate(value string) time.Time {
// 	t, _ := time.Parse(DateISOFormat, value)
// 	return t
// }

// // TimeStartOfDayInHCM ...
// func TimeStartOfDayInHCM(t time.Time) time.Time {
// 	now := time.Now()
// 	l := now.Location()
// 	y, m, d := t.In(l).Date()
// 	return time.Date(y, m, d, 0, 0, 0, 0, l).UTC()
// }

// // TimeEndOfDayHCM ...
// func TimeEndOfDayHCM(t time.Time) time.Time {
// 	//l, _ := time.LoadLocation(timezoneHCM)
// 	now := time.Now()
// 	l := now.Location()
// 	y, m, d := t.In(l).Date()
// 	return time.Date(y, m, d, 23, 59, 59, 0, l).UTC()
// }
