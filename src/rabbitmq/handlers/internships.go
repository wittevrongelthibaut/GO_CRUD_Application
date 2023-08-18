package handlers

import (
	"fmt"
	"log"
	db "main/src/data"
	"strings"
	"main/src/domain"
	idto "main/src/domain/internship"
	amqp "github.com/rabbitmq/amqp091-go"
	recommendation "main/src/domain/recommendation"
)

func HandleInternshipMessage(msg domain.Message, ch *amqp.Channel) interface{} {
	data := msg.Data.(map[string]interface{})
	switch strings.ToLower(msg.Action){
	case "get":

		if len(data) == 0 || data["id"] == nil {
			return db.GetAllInternships()
		} else {
			return db.GetInternshipById(int(data["id"].(float64)))
		}

	case "add":
		response, err := db.AddInternship(
			idto.InternshipWriteDto{ CompanyId: int(data["companyId"].(float64)),
									 ManagerId: int(data["managerId"].(float64)),
									 SupervisorId: int(data["supervisorId"].(float64)),
									 Title: data["title"].(string), Description: data["description"].(string),
									 CategoryId: int(data["categoryId"].(float64)), Tags: data["tags"].(string),
									 Location: data["location"].(string), StartDate: data["startDate"].(string),
									 EndDate: data["endDate"].(string)})

		if err != nil {
			return domain.Error{Error: err.Error()}
		}

		return response

	case "assignstudent":

		response, err := db.AssignStudentToInternship(int(data["internshipId"].(float64)), int(data["studentId"].(float64)))

		if err != nil {
			return domain.Error{Error: err.Error()}
		}

		return response

	case "approve":

		response, err := db.ApproveInternship(int(data["internshipId"].(float64)))

		if err != nil {
			return domain.Error{Error: err.Error()}
		}

		idto.CreateContracts(ch, int(data["internshipId"].(float64)))

		return response
	
	case "deny":
		
		response, err := db.DenyInternship(int(data["internshipId"].(float64)))

		if err != nil {
			return domain.Error{Error: err.Error()}
		}

		return response
	
	case "recommendations":

		category := recommendation.RetrieveRecommendationFromRPC(int(data["id"].(float64)), ch)
		fmt.Print(category)
		return db.GetStudentIntershipsSortedByCategory(category.CategoryId)

	default:
		err := domain.Error{Error: fmt.Sprintf("unkown action: %s", msg.Action)}
		log.Print(err)
		return err
	}
}
