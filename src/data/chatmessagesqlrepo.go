package database

import (
	"context"
	"fmt"
	dto "main/src/domain/chatmessage"
	"main/src/domain"
	"os"
)

func GetAllChatMessages(companyId int, studentId int) []dto.ChatmessageReadDTO {
	dbpool := GetConnection()
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), `SELECT m.id as message_id, c.id as company_id, s.id as student_id, 
													m.content, m.sent_by_student, c.name as company_name,  
													CONCAT(s.firstname, ' ', s.lastname) as student_name,
													m.sent_at
													FROM messages m
													INNER JOIN companies c ON m.company_id = c.id 
													INNER JOIN students s ON m.student_id = s.id 
													WHERE m.company_id = $1 AND m.student_id = $2`, companyId, studentId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	var messages []dto.ChatmessageReadDTO

	for rows.Next() {
		var m dto.ChatmessageReadDTO
		err = rows.Scan(&m.MessageId, &m.CompanyId, &m.StudentId, &m.Content, &m.SentByStudent, &m.CompanyName, &m.StudentName, &m.SentAt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan the row: %v\n", err)
		}
		messages = append(messages, m)
	}

	return messages
}

func AddChatMessage(c dto.ChatmessageWriteDTO) (domain.Created, error) {
	dbpool := GetConnection()
	defer dbpool.Close()

	_, err := dbpool.Exec(context.Background(), "INSERT INTO messages (company_id, student_id, content, sent_by_student) VALUES ($1, $2, $3, $4)", c.CompanyId, c.StudentId, c.Content, c.SentByStudent)
	if err != nil {
		return domain.Created{}, err
	}

	return domain.Created{Confirmation: "Chatmessage created succesfully"}, nil
}
