package handlers

import (
	"fmt"
	"log"
	db "main/src/data"
	"main/src/domain"
	dto "main/src/domain/student"
	amqp "github.com/rabbitmq/amqp091-go"
	"strings"
)

func HandleStudentMessage(msg domain.Message, ch *amqp.Channel) interface{} {
	data := msg.Data.(map[string]interface{})
	switch strings.ToLower(msg.Action){
	case "get":
		if len(data) == 0 || data["id"] == nil {
			return db.GetAllStudents()
		} else {
			return db.GetStudentById(int(data["id"].(float64)))
		}
	case "add":
		response, err := db.AddStudent(dto.StudentWriteDto{Firstname: data["firstname"].(string), Lastname: data["lastname"].(string), Phone: data["phone"].(string)})
		if err != nil {
			return domain.Error{Error: err.Error()}
		}
		return response
	case "studentinternships":
		return db.GetStudentInternship(int(data["id"].(float64)))
	default:
		err := domain.Error{Error: fmt.Sprintf("unkown action: %s", msg.Action)}
		log.Print(err)
		return err
	}
}