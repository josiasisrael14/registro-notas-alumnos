package model

type Teacher struct {
	IdTeacher string `json:"idTeacher"`
	Name      string `json:"name"`
	Surnames  string `json:"surnames"`
	CellPone  string `json:"cellPone"`
}

type Teachers []Teacher

type ResponseStatusTeacher struct {
	Response string `json:"response"`
}
