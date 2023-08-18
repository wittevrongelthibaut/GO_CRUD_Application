package student

type StudentReadDto struct {
	Id int `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Phone *string `json:"phone"`
}