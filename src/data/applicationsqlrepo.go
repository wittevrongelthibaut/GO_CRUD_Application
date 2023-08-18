package database

import (
	"context"
	"fmt"
	dto "main/src/domain/application"
	"main/src/domain"
	"os"
)

func GetAllApplicationsByInternshipId(id int) []dto.ApplicationReadDto {
	dbpool := GetConnection()
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), `select a.*, s.id as student_id, s.firstname as student_name, s.lastname as student_lastname, s.phone as student_phone
												     from applications a
												     inner join students s on a.intern_id = s.id
												     where internship_id = $1`, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	applications := []dto.ApplicationReadDto{}

	for rows.Next() {
		var a dto.ApplicationReadDto
		err = rows.Scan(&a.InternshipId, &a.InternId, &a.Motivation, &a.Accepted, &a.Student.Id, &a.Student.Firstname, &a.Student.Lastname, &a.Student.Phone)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan the row: %v\n", err)
		}
		applications = append(applications, a)
	}
	fmt.Print(applications)
	return applications
}

func AddApplication(a dto.ApplicationWriteDto) (domain.Created, error) {
	dbpool := GetConnection()
	defer dbpool.Close()

	_, err := dbpool.Exec(context.Background(), "INSERT INTO applications (internship_id, intern_id, motivation) VALUES ($1, $2, $3)", a.InternshipId, a.InternId, a.Motivation)
	if err != nil {
		return domain.Created{}, err
	}

	return domain.Created{Confirmation: "Application created succesfully"}, nil
}

func UpdateApplication(a dto.ApplicationUpdateDto) (domain.Created, error) {
	dbpool := GetConnection()
	defer dbpool.Close()

	_, err := dbpool.Exec(context.Background(), "UPDATE applications SET accepted = $1 WHERE internship_id = $2 AND intern_id = $3", a.Accepted, a.InternshipId, a.InternId)
	if err != nil {
		return domain.Created{}, err
	}

	return domain.Created{Confirmation: "Application updated succesfully"}, nil
}

