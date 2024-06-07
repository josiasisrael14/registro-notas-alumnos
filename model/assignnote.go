package model

import "time"

type AssignNote struct {
	Note_id            string    `json:"note_id"`
	NameStudent        string    `json:"nameStudent"`
	NameTeacher        string    `json:"nameTeacher"`
	NameStuff          string    `json:"nameStuff"`
	Grade              string    `json:"grade"`
	SpecificationLevel string    `json:"specificationLevel"`
	Section            string    `json:"section"`
	Note               string    `json:"note"`
	Date               time.Time `json:"date"`
	Comments           string    `json:"comments"`
}

type AssignNotes []AssignNote

type ResponseStatusAssignNote struct {
	Response string `json:"response"`
}
