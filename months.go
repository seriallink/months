package timex

import "time"

func Months(t time.Time) int {
	return t.Year() * 12 + int(t.Month())
}

// Calculates the difference between two times in months
func MonthsBetween(t1, t2 time.Time) (int) {
	return Months(t2) - Months(t1)
}

func FirstDayOfTheMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func LastDayOfTheMonth(t time.Time) time.Time {
	return FirstDayOfTheMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func AddMonths(t time.Time, n int) (r time.Time) {

	// add months
	r = time.Date(t.Year(), t.Month() + time.Month(n), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	// check difference between months
	if MonthsBetween(t,r) != n {
		for {
			r = r.Add(-time.Hour*24)
			if MonthsBetween(t,r) == n {
				break
			}
		}
	}

	return
}
