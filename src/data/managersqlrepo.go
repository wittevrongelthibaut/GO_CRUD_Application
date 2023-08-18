package database

import (
	"context"
	"fmt"
	dto "main/src/domain/manager"
	"os"
	"main/src/domain"
)

func GetAllManagers() []dto.ManagerReadDTO {
	dbpool := GetConnection()
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), "SELECT * FROM managers")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v \n", err)
	}

	var managers []dto.ManagerReadDTO

	for rows.Next() {
		var m dto.ManagerReadDTO
		err = rows.Scan(&m.Id, &m.Firstname, &m.Lastname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan the row: %v \n", err)
		}
		managers = append(managers, m)
	}

	return managers
}

func GetManagerById(id int) dto.ManagerReadDTO {
	dbpool := GetConnection()
	defer dbpool.Close()

	var m dto.ManagerReadDTO
	err := dbpool.QueryRow(context.Background(), "SELECT * FROM managers WHERE id = $1", id).Scan(&m.Id, &m.Firstname, &m.Lastname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v \n", err)
	}

	return m
}

func AddManager(m dto.ManagerWriteDTO) (domain.Created, error) {
	dbpool := GetConnection()
	defer dbpool.Close()

	_, err := dbpool.Exec(context.Background(), "INSERT INTO managers (firstname, lastname) VALUES ($1, $2)", m.Firstname, m.Lastname)
	if err != nil {
		return domain.Created{}, err
	}

	return domain.Created{Confirmation: "Manager created succesfully"}, nil
}


