package database

import (
	"context"
	"fmt"
	"os"
	dto "main/src/domain/category"
)

func GetAllCategories() []dto.CategoryReadDTO {
	dbpool := GetConnection()
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), "SELECT * FROM categories")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	var categories []dto.CategoryReadDTO

	for rows.Next() {
		var c dto.CategoryReadDTO
		err = rows.Scan(&c.Id, &c.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan the row: %v\n", err)
		}
		categories = append(categories, c)
	}

	return categories
}