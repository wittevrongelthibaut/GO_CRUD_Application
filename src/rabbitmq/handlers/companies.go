package handlers

import (
	"fmt"
	"log"
	"main/src/domain"
	"strings"
	db "main/src/data"
	cdto "main/src/domain/company"
) 

func HandleCompanyMessage(msg domain.Message) interface{} {
	data := msg.Data.(map[string]interface{})
	switch strings.ToLower(msg.Action){
	case "get":
		if len(data) == 0 || data["id"] == nil {
			return db.GetAllCompanies()
		} else {
			return db.GetCompanyById(int(data["id"].(float64)))
		}
	case "add":
		response, err := db.AddCompany(cdto.Company{Name: data["name"].(string)})
		if err != nil {
			return domain.Error{Error: err.Error()}
		}
		return response
	case "getinternships":
		return db.GetCompanyInternships(int(data["id"].(float64)))	
	default:
		err := domain.Error{Error: fmt.Sprintf("unkown action: %s", msg.Action)}
		log.Print(err)
		return err
	}
}