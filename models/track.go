package models

type Track struct {
	ID       int32  `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Singer   string `json:"singer" db:"singer"`
	Album    string `json:"album" db:"album"`
	Duration string `json:"duration" db:"duration"`
}
