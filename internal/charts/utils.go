package charts

import "time"

func getHourlyTimeRange() []time.Time {
	var times []time.Time
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	for i := 0; i < 24; i++ {
		times = append(times, startOfDay.Add(time.Duration(i)*time.Hour))
	}

	return times
}

func isLeapYear(year int16) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func daysInMonth(month int8, year int16) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 0
	}
}
