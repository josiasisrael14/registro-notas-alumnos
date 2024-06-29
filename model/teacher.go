package model

type Teacher struct {
	IdTeacher   string `json:"idTeacher"`
	Name        string `json:"name"`
	Surnames    string `json:"surnames"`
	CellPone    string `json:"cellPone"`
	UserTeacher string `json:"userTeacher,omitempty"`
	Password    string `json:"password,omitempty"`
	RolId       string `json:"rolId,omitempty"`
}

type Teachers []Teacher

type ResponseStatusTeacher struct {
	Response string `json:"response"`
}
