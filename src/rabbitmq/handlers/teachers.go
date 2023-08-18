package handlers

import (
	"fmt"
	"log"
	"main/src/domain"
	"strings"
	db "main/src/data"
	dto "main/src/domain/teacher"
)

func HandleTeacherMessage(msg domain.Message) interface{} {
	data := msg.Data.(map[string]interface{})
	switch strings.ToLower(msg.Action){
	case "get":
		if len(data) == 0 || data["id"] == nil {
			return db.GetAllTeachers()
		} else {
			return db.GetTeacherById(int(data["id"].(float64)))
		}
	case "add":
		response, err := db.AddTeacher(dto.TeacherWriteDTO{Firstname: data["firstname"].(string), Lastname: data["lastname"].(string)})
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
