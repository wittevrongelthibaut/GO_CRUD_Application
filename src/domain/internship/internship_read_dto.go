package internship

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Internship struct {
	Id int `json:"id"`
	CompanyId int `json:"companyId"`
	InternId pgtype.Int8 `json:"internId"`
	ManagerId int `json:"managerId"`
	SupervisorId int `json:"supervisorId"`
	Title string `json:"title"`
	Description string `json:"description"`
	CategoryId int `json:"categoryId"`
	Tags string `json:"tags"`
	Location string `json:"location"`
	Approved pgtype.Bool `json:"approved"`
	StartDate pgtype.Date `json:"startDate"`
	EndDate pgtype.Date `json:"endDate"`
	CompanyContract pgtype.Int8 `json:"companyContract"`
	StudentContract pgtype.Int8 `json:"internContract"`
	ManagerContract pgtype.Int8 `json:"managerContract"`
	CompanyName string `json:"companyName"`
	StudentFirstname *string `json:"internFirstname"`
	StudentLastname *string `json:"internLastname"`
	ManagerFirstname string `json:"managerFirstname"`
	ManagerLastname string `json:"managerLastname"`
	SupervisorFirstname string `json:"supervisorFirstname"`
	SupervisorLastname string `json:"supervisorLastname"`
	CategoryName string `json:"categoryName"`
}