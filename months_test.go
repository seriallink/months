package timex

import (
	"testing"
	"time"
)

var between = []struct {
	t1 time.Time
	t2 time.Time
	n int
}{
	{
		time.Date(2017, time.Month(1), 10, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.Month(1), 10, 0, 0, 0, 0, time.Local),
		0,
	},
	{
		time.Date(2017, time.Month(1), 10, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.Month(1), 10, 7, 0, 0, 0, time.Local),
		0,
	},
	{
		time.Date(2017, time.Month(1), 10, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.Month(2), 10, 0, 0, 0, 0, time.Local),
		1,
	},
	{
		time.Date(2017, time.Month(1), 31, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.Month(2), 28, 0, 0, 0, 0, time.Local),
		1,
	},
	{
		time.Date(2017, time.Month(1), 1, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.Month(3), 31, 0, 0, 0, 0, time.Local),
		2,
	},
	{
		time.Date(2017, time.Month(11), 10, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.Month(10), 10, 0, 0, 0, 0, time.Local),
		-1,
	},
	{
		time.Date(2017, time.Month(1), 10, 0, 0, 0, 0, time.Local),
		time.Date(2016, time.Month(2), 10, 0, 0, 0, 0, time.Local),
		13,
	},
	{
		time.Date(2017, time.Month(9), 10, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.Month(9), 10, 0, 0, 0, 1, time.Local),
		1,
	},
}

var add = []struct {
	t1 time.Time
	t2 time.Time
	n int
}{
	{
		time.Date(2017, time.Month(1), 10, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.Month(2), 10, 0, 0, 0, 0, time.Local),
		1,
	},
	{
		time.Date(2017, time.Month(1), 31, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.Month(2), 28, 0, 0, 0, 0, time.Local),
		1,
	},
	{
		time.Date(2020, time.Month(1), 31, 0, 0, 0, 0, time.Local),
		time.Date(2020, time.Month(2), 29, 0, 0, 0, 0, time.Local),
		1,
	},
	{
		time.Date(2017, time.Month(12), 31, 0, 0, 0, 0, time.Local),
		time.Date(2018, time.Month(2), 28, 0, 0, 0, 0, time.Local),
		2,
	},
	{
		time.Date(2017, time.Month(1), 10, 0, 0, 0, 0, time.Local),
		time.Date(2018, time.Month(1), 10, 0, 0, 0, 0, time.Local),
		12,
	},
	{
		time.Date(2017, time.Month(8), 10, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.Month(6), 10, 0, 0, 0, 0, time.Local),
		-2,
	},
	{
		time.Date(2017, time.Month(12), 31, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.Month(11), 30, 0, 0, 0, 0, time.Local),
		-1,
	},
}

func TestMonthsBetween(t *testing.T) {
	for _, d := range between {
		if MonthsBetween(d.t1, d.t2) != d.n {
			t.Fail()
		}
	}
}

func TestAddMonths(t *testing.T) {
	for _, d := range add {
		if AddMonths(d.t1, d.n) != d.t2 {
			t.Fail()
		}
	}
}
