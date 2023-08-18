package database

import (
	"context"
	"fmt"
	"os"
	dto "main/src/domain/student"
	idto "main/src/domain/internship"
	"main/src/domain"
)

func GetAllStudents() []dto.StudentReadDto {
	dbpool := GetConnection()
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), "SELECT * FROM students")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	var students []dto.StudentReadDto

	for rows.Next() {
		var s dto.StudentReadDto
		err = rows.Scan(&s.Id, &s.Firstname, &s.Lastname, &s.Phone)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan the row: %v\n", err)
		}
		students = append(students, s)
	}

	return students
}

func GetStudentById(id int) dto.StudentReadDto {
	dbpool := GetConnection()
	defer dbpool.Close()

	var s dto.StudentReadDto
	err := dbpool.QueryRow(context.Background(), "SELECT * FROM students WHERE id = $1", id).Scan(&s.Id, &s.Firstname, &s.Lastname, &s.Phone)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return s
}

func AddStudent(s dto.StudentWriteDto) (domain.Created, error) {
    dbpool := GetConnection()
    defer dbpool.Close()

    _, err := dbpool.Exec(context.Background(), "INSERT INTO students (firstname, lastname, phone) VALUES ($1, $2, $3)", s.Firstname, s.Lastname, s.Phone)
    if err != nil {
        return domain.Created{}, err
    }

    return domain.Created{Confirmation: "Student created succesfully"}, nil
}

func GetStudentInternship(id int) idto.Internship {
	dbpool := GetConnection()
	defer dbpool.Close()

	var i idto.Internship
	err := dbpool.QueryRow(context.Background(), `SELECT i.*, c.name as company_name, s.firstname as student_firstname, s.lastname as student_lastname,m.firstname 
												  as manager_firstname, m.lastname as manager_lastname, l.firstname as lecturer_firstname, l.lastname as lecturer_lastname, cat.name as category_name
												  FROM internships i
												  INNER JOIN companies c ON i.company_id = c.id 
												  LEFT JOIN students s ON i.intern_id = s.id
												  INNER JOIN managers m ON i.manager_id = m.id
												  INNER JOIN categories cat ON i.category_id = cat.id
												  INNER JOIN lecturers l ON i.supervisor_id = l.id
												  WHERE intern_id = $1`, id).Scan(&i.Id, &i.CompanyId, &i.InternId, &i.ManagerId, &i.SupervisorId, &i.Title, &i.Description, &i.CategoryId, &i.Tags, &i.Location, &i.Approved, &i.StartDate, &i.EndDate, &i.CompanyContract, &i.StudentContract, &i.ManagerContract, &i.CompanyName, &i.StudentFirstname, &i.StudentLastname, &i.ManagerFirstname, &i.ManagerLastname, &i.SupervisorFirstname, &i.SupervisorLastname, &i.CategoryName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return i
}
