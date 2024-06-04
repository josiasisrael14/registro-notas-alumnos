package model

type Degree struct {
	IdDegree      string `json:"idDegree"`
	NameDegree    string `json:"nameDegree"`
	SpecificLevel string `json:"specificLevel"`
}

type Degrees []Degree

type ResponseStatusDegree struct {
	Response string `json:"response"`
}
