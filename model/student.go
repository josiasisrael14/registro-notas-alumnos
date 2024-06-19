package model

import (
	"time"
)

type CustomDate time.Time

type Student struct {
	StudentId   string    `json:"studentId"`
	NameStudent string    `json:"nameStudent"`
	LastName    string    `json:"lastName"`
	BirthDate   time.Time `json:"birthDate"`
}

type Students []Student

type ResponseStatusStudent struct {
	Response string `json:"response"`
}

/*// truncateTime elimina la parte del tiempo de un objeto time.Time
func truncateTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func ConvertTime(students Students) Students {

	for i, studentResult := range students {
		students[i].BirthDate = truncateTime(studentResult.BirthDate)
	}

	return students
}*/
