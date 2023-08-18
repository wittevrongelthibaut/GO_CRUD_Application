package database

import (
	"context"
	"fmt"
	"os"
	"main/src/domain"
	idto "main/src/domain/internship"
)

func GetAllInternships() []idto.Internship {
	dbpool := GetConnection()
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), `SELECT i.*, c.name as company_name, s.firstname as student_firstname, s.lastname as student_lastname,m.firstname 
													 as manager_firstname, m.lastname as manager_lastname, l.firstname as lecturer_firstname, l.lastname as lecturer_lastname, cat.name as category_name
													 FROM internships i
													 INNER JOIN companies c ON i.company_id = c.id 
													 LEFT JOIN students s ON i.intern_id = s.id
													 INNER JOIN managers m ON i.manager_id = m.id
													 INNER JOIN categories cat ON i.category_id = cat.id
													 INNER JOIN lecturers l ON i.supervisor_id = l.id`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	var internships []idto.Internship

	for rows.Next() {
		var i idto.Internship
		err = rows.Scan(&i.Id, &i.CompanyId, &i.InternId, &i.ManagerId, &i.SupervisorId, &i.Title, &i.Description, &i.CategoryId, &i.Tags, &i.Location, &i.Approved, &i.StartDate, &i.EndDate, &i.CompanyContract, &i.StudentContract, &i.ManagerContract, &i.CompanyName, &i.StudentFirstname, &i.StudentLastname, &i.ManagerFirstname, &i.ManagerLastname, &i.SupervisorFirstname, &i.SupervisorLastname, &i.CategoryName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan the row: %v\n", err)
		}
		internships = append(internships, i)
	}

	return internships
}

func GetInternshipById(id int) idto.Internship {
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
												  WHERE i.id = $1`, id).Scan(&i.Id, &i.CompanyId, &i.InternId, &i.ManagerId, &i.SupervisorId, &i.Title, &i.Description, &i.CategoryId, &i.Tags, &i.Location, &i.Approved, &i.StartDate, &i.EndDate, &i.CompanyContract, &i.StudentContract, &i.ManagerContract, &i.CompanyName, &i.StudentFirstname, &i.StudentLastname, &i.ManagerFirstname, &i.ManagerLastname, &i.SupervisorFirstname, &i.SupervisorLastname, &i.CategoryName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return i
}

func AddInternship(i idto.InternshipWriteDto) (domain.Created, error) {
	fmt.Println(i.CompanyId)
	dbpool := GetConnection()
	defer dbpool.Close()

	_, err := dbpool.Exec(context.Background(), 
	"INSERT INTO internships (company_id, manager_id, supervisor_id, title, description, category_id, tags, location, start_date, end_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", i.CompanyId, i.ManagerId, i.SupervisorId, i.Title, i.Description, i.CategoryId, i.Tags, i.Location, i.StartDate, i.EndDate)
	if err != nil {
		return domain.Created{}, err
	}

	return domain.Created{Confirmation: "Internship created succesfully"}, nil
}

func AssignStudentToInternship(internshipId int, studentId int) (domain.Created, error) {
	dbpool := GetConnection()
	defer dbpool.Close()

	_, err := dbpool.Exec(context.Background(), "UPDATE internships SET intern_id = $1 WHERE id = $2", studentId, internshipId)
	if err != nil {
		return domain.Created{}, err
	}

	return domain.Created{Confirmation: "Student assigned to internship succesfully"}, nil
}

func ApproveInternship(internshipId int) (domain.Created, error) {
	dbpool := GetConnection()
	defer dbpool.Close()

	_, err := dbpool.Exec(context.Background(), "UPDATE internships SET approved = true WHERE id = $1", internshipId)
	if err != nil {
		return domain.Created{}, err
	}

	return domain.Created{Confirmation: "Internship approved succesfully"}, nil
}

func DenyInternship (internshipId int) (domain.Created, error) {
	dbpool := GetConnection()
	defer dbpool.Close()

	_, err := dbpool.Exec(context.Background(), "UPDATE internships SET approved = false WHERE id = $1", internshipId)
	if err != nil {
		return domain.Created{}, err
	}

	return domain.Created{Confirmation: "Internship denied succesfully"}, nil
}

func GetStudentIntershipsSortedByCategory(categoryId int) []idto.Internship {
	dbpool := GetConnection()
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), `SELECT i.*, c.name as company_name, s.firstname as student_firstname, s.lastname as student_lastname,m.firstname 
	as manager_firstname, m.lastname as manager_lastname, l.firstname as lecturer_firstname, l.lastname as lecturer_lastname, cat.name as category_name
	FROM internships i
	INNER JOIN companies c ON i.company_id = c.id 
	LEFT JOIN students s ON i.intern_id = s.id
	INNER JOIN managers m ON i.manager_id = m.id
	INNER JOIN categories cat ON i.category_id = cat.id
	INNER JOIN lecturers l ON i.supervisor_id = l.id
	ORDER BY 
	CASE WHEN category_id = $1 THEN 1 ELSE 2 END, category_id ASC`, categoryId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	var internships []idto.Internship

	for rows.Next() {
		var i idto.Internship
		err = rows.Scan(&i.Id, &i.CompanyId, &i.InternId, &i.ManagerId, &i.SupervisorId, &i.Title, &i.Description, &i.CategoryId, &i.Tags, &i.Location, &i.Approved, &i.StartDate, &i.EndDate, &i.CompanyContract, &i.StudentContract, &i.ManagerContract, &i.CompanyName, &i.StudentFirstname, &i.StudentLastname, &i.ManagerFirstname, &i.ManagerLastname, &i.SupervisorFirstname, &i.SupervisorLastname, &i.CategoryName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan the row: %v\n", err)
		}
		internships = append(internships, i)
	}

	return internships
}

