package internship

type InternshipWriteDto struct {
	CompanyId int `json:"companyId"`
	ManagerId int `json:"managerIDd"`
	SupervisorId int `json:"supervisorId"`
	Title string `json:"title"`
	Description string `json:"description"`
	CategoryId int `json:"categoryId"`
	Tags string `json:"tags"`
	Location string `json:"location"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
}