package timex

import "time"

func WeeksBreakdown(t1, t2 time.Time) (dates []time.Time) {

	dates = append(dates, t1)

	for t1.Before(t2) {

		t1 = t1.Add(time.Hour * 24)

		if t1.Weekday() == 0 {
			dates = append(dates, t1)
		}

	}

	return dates

}
