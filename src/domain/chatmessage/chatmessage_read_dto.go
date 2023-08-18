package chatmessage

import "github.com/jackc/pgx/v5/pgtype"

/*type ChatmessageReadDTO struct {
	Id int `json:"id"`
	CompanyId int `json:"companyId"`
	StudentId int `json:"studentId"`
	Content string `json:"content"`
	SentByStudent bool `json:"sentByStudent"`
	SentAt pgtype.Timestamp `json:"sentAt"`
} */

type ChatmessageReadDTO struct {
	MessageId int `json:"messageId"`
	CompanyId int `json:"companyId"`
	StudentId int `json:"studentId"`
	Content string `json:"content"`
	SentByStudent bool `json:"sentByStudent"`
	CompanyName string `json:"companyName"`
	CompanyEmail string `json:"companyEmail"`
	StudentEmail string `json:"studentEmail"`
	StudentName string `json:"studentName"`
	SentAt pgtype.Timestamp `json:"sentAt"`
}
