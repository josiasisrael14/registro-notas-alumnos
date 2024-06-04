package model

type Section struct {
	IdSection   string `json:"idSection"`
	NameSection string `json:"nameSection"`
}

type Sections []Section

type ResponseStatusSection struct {
	Response string `json:"response"`
}
