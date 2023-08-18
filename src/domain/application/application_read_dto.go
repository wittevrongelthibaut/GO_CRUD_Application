package application

import (
	"main/src/domain/student"

	"github.com/jackc/pgx/v5/pgtype"
)

type ApplicationReadDto struct {
	InternshipId int `json:"internshipId"`
	InternId int `json:"internId"`
	Motivation string `json:"motivation"`
	Accepted pgtype.Bool `json:"accepted"`
	Student student.StudentReadDto `json:"student"`
}