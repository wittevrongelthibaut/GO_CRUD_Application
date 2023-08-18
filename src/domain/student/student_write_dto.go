package student

type StudentWriteDto struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Phone string `json:"phone"`
}