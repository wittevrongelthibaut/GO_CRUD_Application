package database

import (
	"context"
	"fmt"
	dto "main/src/domain/teacher"
	"os"
	"main/src/domain"
)

func GetAllTeachers() []dto.TeacherReadDTO {
	dbpool := GetConnection()
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), "SELECT * FROM lecturers")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v \n", err)
	}

	var teachers []dto.TeacherReadDTO

	for rows.Next() {
		var t dto.TeacherReadDTO
		err = rows.Scan(&t.Id, &t.Firstname, &t.Lastname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan the row: %v \n", err)
		}
		teachers = append(teachers, t)
	}

	return teachers
}

func GetTeacherById(id int) dto.TeacherReadDTO {
	dbpool := GetConnection()
	defer dbpool.Close()

	var t dto.TeacherReadDTO
	err := dbpool.QueryRow(context.Background(), "SELECT * FROM lecturers WHERE id = $1", id).Scan(&t.Id, &t.Firstname, &t.Lastname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v \n", err)
	}

	return t
}

func AddTeacher(t dto.TeacherWriteDTO) (domain.Created, error) {
	dbpool := GetConnection()
	defer dbpool.Close()

	_, err := dbpool.Exec(context.Background(), "INSERT INTO lecturers (firstname, lastname) VALUES ($1, $2)", t.Firstname, t.Lastname)
	if err != nil {
		return domain.Created{}, err
	}

	return domain.Created{Confirmation: "Teacher created succesfully"}, nil
}