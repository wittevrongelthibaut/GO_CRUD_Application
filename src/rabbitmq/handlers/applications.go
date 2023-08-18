package handlers

import (
	"main/src/domain"
	"strings"
	"fmt"
	"log"
	db "main/src/data"
	dto "main/src/domain/application"
)

func HandleApplicationMessage(msg domain.Message) interface{} {
	data := msg.Data.(map[string]interface{})
	switch strings.ToLower(msg.Action){
	case "get":
		return db.GetAllApplicationsByInternshipId(int(data["id"].(float64)))
	case "add":
		response, err := db.AddApplication(dto.ApplicationWriteDto{ InternshipId: int(data["internshipId"].(float64)), InternId: int(data["internId"].(float64)), Motivation: data["motivation"].(string)})
		if err != nil {
			return domain.Error{Error: err.Error()}
		}
		return response
	case "update":
		response, err := db.UpdateApplication(dto.ApplicationUpdateDto{ InternshipId: int(data["internshipId"].(float64)), InternId: int(data["internId"].(float64)), Accepted: data["accepted"].(bool)})
		if err != nil {
			return domain.Error{Error: err.Error()}
		}
		if data["accepted"].(bool) {
			db.AssignStudentToInternship(int(data["internshipId"].(float64)), int(data["internId"].(float64)))
		}
		return response
	default:
		err := domain.Error{Error: fmt.Sprintf("unkown action: %s", msg.Action)}
		log.Print(err)
		return err
	}
}