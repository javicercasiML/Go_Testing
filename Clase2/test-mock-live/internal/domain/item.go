package domain

import "time"

type Item struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Weight  float64   `json:"weight"`
	Price   float64   `json:"price"`
	Release time.Time `json:"release"`
}
func (i *Item) Valid() bool {
	return i.Release.Before(time.Now())
}
