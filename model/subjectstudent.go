package model

type StudentSubject struct {
	IdSubjectStudent   string `json:"idSubjectStudent,omitempty"`
	IdStudent          string `json:"idStudent,omitempty"`
	NameStudent        string `json:"nameStudent,omitempty"`
	LastName           string `json:"lastName,omitempty"`
	Grade              string `json:"grade"`
	SpecificationLevel string `json:"specificationLevel,omitempty"`
	Section            string `json:"section"`
}

type StudentSubjects []StudentSubject

type ResponseStatusSubjectStudent struct {
	Response string `json:"response"`
}
