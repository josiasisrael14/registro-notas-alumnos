package model

type SubjectTeacher struct {
	IdSubjectTeacher string `json:"idSubjectTeacher"`
	NameTeacher      string `json:"nameTeacher"`
	Surnames         string `json:"surnames"`
	Grade            string `json:"grade"`
	Section          string `json:"section"`
	Subject          string `json:"subject"`
	Comments         string `json:"comments"`
}

type SubjectTeachers []SubjectTeacher

type ResponseStatusSubjectTeacher struct {
	Response string `json:"response"`
}
