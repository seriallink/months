package timex

import (
	"time"
)

// Returns date with the time portion of the day truncated
func Trunc(t time.Time) (time.Time) {
	_, offset := t.Zone()
	seconds := t.Unix()
	return time.Unix(seconds-(seconds+int64(offset))%86400, 0)
}

// Returns today's time to midnight
func Today() (time.Time) {
	return Trunc(time.Now())
}

// Returns tomorrow's time by adding one day from today
func Tomorrow() (time.Time) {
	return AddDays(Today(), 1)
}

// Returns tomorrow's time by subtracting one day from today
func Yesterday() (time.Time) {
	return SubDays(Today(), 1)
}

func Days(t time.Time) int {
	_, offset := t.Zone()
	return int((t.Unix() + int64(offset)) / int64(86400))
}

// Calculates the difference between two times in days
func DaysBetween(t1, t2 time.Time) (int) {
	return Days(t2) - Days(t1)
}

// Adds "n" days to a given time
func AddDays(t time.Time, n int) (ret time.Time) {
	ret = Trunc(t)
	if n != 0 {
		ret = time.Unix(ret.Unix() + int64(86400*n), 0)
	}
	return
}

// Subtracts "n" days to a given time
func SubDays(t time.Time, n int) (ret time.Time) {
	ret = Trunc(t)
	if n != 0 {
		ret = time.Unix(ret.Unix() - int64(86400*n), 0)
	}
	return
}

func DaysBreakdown(t1, t2 time.Time) (dates []time.Time) {

	dates = append(dates,t1)
	count := DaysBetween(t2,t1)

	for i:=0; i<count; i++ {
		t1 = t1.Add(time.Hour * 24)
		dates = append(dates,t1)
	}

	return

}
