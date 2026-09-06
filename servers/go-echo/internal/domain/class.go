package domain

import "time"

type Period struct {
	DayOfWeek DayOfWeek
	StartTime int
	EndTime   int
	Room      Room
}

type Class struct {
	ID         string
	Name       string
	Department Department
	Teachers   []Teacher
	Periods    []Period
	Semesters  []Semester
	Year       int
	Credit     int
}

type ClassSchedule struct {
	ClassID     string
	Count       int
	Description string
}

type Assignment struct {
	ID            string
	ClassSchedule ClassSchedule
	Title         string
	Description   string
	MaxScore      int
	DueDate       time.Time
	Attachments   []AssignmentAttachment
}

type AssignmentAttachment struct {
	Name string
	Url  string
}
