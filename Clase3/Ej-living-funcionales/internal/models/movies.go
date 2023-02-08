package models

type Movie struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Rating float64 `json:"rating"`
	Year   int     `json:"year"`
}

func (m *Movie) Valid() bool {
	return m.Title != "" && m.Rating >= 0 && m.Rating <= 10 && m.Year >= 0
}