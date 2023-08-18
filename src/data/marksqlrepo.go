package database

import (
	"context"
	"fmt"
	"os"
	dto "main/src/domain/marks"
)

func GetMarkFromStudent(studentId int) dto.MarkReadDto {
	dbpool := GetConnection()
	defer dbpool.Close()

	var i dto.MarkReadDto
	err := dbpool.QueryRow(context.Background(), "SELECT * FROM marks WHERE student_id = $1", studentId).Scan(&i.Id, &i.StudentId, &i.Oop, &i.Linux, &i.Win, &i.Cms, &i.Sec, &i.Dotnet, &i.Laravel, &i.Db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return i
}