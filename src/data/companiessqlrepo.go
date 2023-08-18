package database

import (
	"context"
	"fmt"
	cdto "main/src/domain/company"
	idto "main/src/domain/internship"
	"main/src/domain"
	"os"
)

func GetAllCompanies() []cdto.Company{
	dbpool := GetConnection()
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), "SELECT * FROM companies")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	var companies []cdto.Company

	for rows.Next() {
		var c cdto.Company
		err = rows.Scan(&c.Id, &c.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan the row: %v\n", err)
		}
		companies = append(companies, c)
	}

	return companies
}

func GetCompanyById(id int) cdto.Company {
	dbpool := GetConnection()
	defer dbpool.Close()

	var c cdto.Company
	err := dbpool.QueryRow(context.Background(), "SELECT * FROM companies WHERE id = $1", id).Scan(&c.Id, &c.Name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return c
}

func AddCompany(c cdto.Company) (domain.Created, error) {
	dbpool := GetConnection()
	defer dbpool.Close()

	_, err := dbpool.Exec(context.Background(), "INSERT INTO companies (name) VALUES ($1)", c.Name)
	if err != nil {
		return domain.Created{}, err
	}

	return domain.Created{Confirmation: "Company created succesfully"}, nil
}

func GetCompanyInternships(id int) []idto.Internship {
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
													 WHERE company_id = $1`, id)
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