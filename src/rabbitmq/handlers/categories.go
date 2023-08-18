package handlers

import (
	"fmt"
	"log"
	"main/src/domain"
	"strings"
	db "main/src/data"
)


func HandleCategoryMessage(msg domain.Message) interface{} {
	switch strings.ToLower(msg.Action){
	case "get":
		return db.GetAllCategories()
	default:
		err := domain.Error{Error: fmt.Sprintf("unkown action: %s", msg.Action)}
		log.Print(err)
		return err
	}
}