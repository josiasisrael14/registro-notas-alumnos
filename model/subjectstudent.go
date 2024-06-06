package model

type StudentSubject struct {
	IdSubjectStudent   string `json:"idSubjectStudent"`
	NameStudent        string `json:"nameStudent"`
	Grade              string `json:"grade"`
	SpecificationLevel string `json:"specificationLevel"`
	Section            string `json:"section"`
}

type StudentSubjects []StudentSubject

type ResponseStatusSubjectStudent struct {
	Response string `json:"response"`
}
