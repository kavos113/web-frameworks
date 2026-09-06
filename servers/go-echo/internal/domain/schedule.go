package domain

type DayOfWeek int

const (
	DayOfWeekMonday DayOfWeek = iota
	DayOfWeekTuesday
	DayOfWeekWednesday
	DayOfWeekThursday
	DayOfWeekFriday
	DayOfWeekSaturday
	DayOfWeekSunday
)

type Semester int

const (
	SemesterSpring Semester = iota
	SemesterSummer
	SemesterFall
	SemesterWinter
)
