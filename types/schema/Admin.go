package schema

import "time"

type Admin struct {
	ID            int64     `json:"id"`
	URL           string    `json:"url"`
	CardSelector  string    `json:"cardSelector"`
	InnerSelector string    `json:"innerSelector"`
	Tag           string    `json:"tag"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
