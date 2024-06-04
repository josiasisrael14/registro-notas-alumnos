package model

type Stuff struct {
	IdStuff     string `json:"idStuff"`
	NameStuff   string `json:"nameStuff"`
	Description string `json:"description"`
}

type Stuffs []Stuff

type ResponseStatusStuff struct {
	Response string `json:"response"`
}
