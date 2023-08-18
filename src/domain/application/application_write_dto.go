package application

type ApplicationWriteDto struct {
	InternshipId int `json:"internshipId"`
	InternId int `json:"internId"`
	Motivation string `json:"motivation"`
}