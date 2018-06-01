# Additional time functions

## Today() time.Time

Returns today's time to midnight

## Tomorrow(t time.Time) time.Time

Returns tomorrow's time by adding one day to the given time

## Yesterday(t time.Time) time.Time

Returns tomorrow's time by subtracting one day from the given time

## DaysBetween(t1, t2 time.Time) (days int)

Calculates the difference between two times in days

## AddDays(t time.Time, n int) time.Time

Adds "n" days to a given time

## SubDays(t time.Time, n int) time.Time

Subtracts "n" days to a given time

## MonthsBetween(t1, t2 time.Time) (days int)

Calculates the difference between two times in months

## AddMonths(t time.Time, n int) time.Time

Adds "n" months to a given time