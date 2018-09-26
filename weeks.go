package timex

import "time"

func LastDayOfWeek(t time.Time, weekday time.Weekday) (day time.Time) {

	day = t

	for {

		if day.Weekday() == weekday {
			break
		}

		day = SubDays(day,1)

	}

	return

}

func NextDayOfWeek(t time.Time, weekday time.Weekday) (day time.Time) {

	day = t

	for {

		if day.Weekday() == weekday {
			break
		}

		day = AddDays(day,1)

	}

	return

}

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
