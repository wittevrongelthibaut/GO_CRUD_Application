package chatmessage

type ChatmessageWriteDTO struct {
	CompanyId int `json:"companyId"`
	StudentId int `json:"studentId"`
	Content string `json:"content"`
	SentByStudent bool `json:"sentByStudent"`
}