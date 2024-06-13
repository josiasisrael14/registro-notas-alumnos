package model

import (
	"time"
)

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
