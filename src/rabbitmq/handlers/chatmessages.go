package handlers

import (
	"main/src/domain"
	"strings"
	db "main/src/data"
	dto "main/src/domain/chatmessage"
	"fmt"
	"log"
)

func HandleChatMessage(msg domain.Message) interface{} {
	data := msg.Data.(map[string]interface{})
	switch strings.ToLower(msg.Action){
	case "get":
		return db.GetAllChatMessages(int(data["companyId"].(float64)), int(data["studentId"].(float64)))
	case "add":
		response, err := db.AddChatMessage(dto.ChatmessageWriteDTO{CompanyId: int(data["companyId"].(float64)), StudentId: int(data["studentId"].(float64)), Content: data["content"].(string), SentByStudent: data["sentByStudent"].(bool)})
		if err != nil {
			return domain.Error{Error: err.Error()}
		}
		return response
	default:
		err := domain.Error{Error: fmt.Sprintf("unkown action: %s", msg.Action)}
		log.Print(err)
		return err
	}
}
