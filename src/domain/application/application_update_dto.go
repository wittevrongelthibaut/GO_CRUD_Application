package application

type ApplicationUpdateDto struct {
	InternshipId int `json:"internshipId"`
	InternId int `json:"internId"`
	Accepted bool `json:"accepted"`
}