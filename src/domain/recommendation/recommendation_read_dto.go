package recommendation

type RecommendationReadDto struct {
	CategoryId int `json:"categoryId"`
	Stars float64 `json:"stars"`
}