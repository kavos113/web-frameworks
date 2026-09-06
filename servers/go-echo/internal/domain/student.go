package domain

import "time"

type Student struct {
	ID         string
	Name       string
	Department Department
	Grade      Grade
}

type Grade int

const (
	GradeUnspecified Grade = iota
	Grade1
	Grade2
	Grade3
	Grade4
	Grade5
	Grade6
	Grade7
	Grade8
	Grade9
)

type StudentClassInfo struct {
	StudentID string
	ClassID   string
	Score     int
}

type StudentAssignmentInfo struct {
	StudentID    string
	AssignmentID string
	Score        int
	SubmittedAt  time.Time
	Attachments  []AssignmentAttachment
}

type Grades struct {
	StudentID     string
	Year          int
	Semester      Semester
	ClassInfos    []StudentClassInfo
	GPA           float64
	GPT           float64
	CreditsEarned int
}
