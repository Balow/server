package models

import (
	"time"
)

// News : todo
type News struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Category string    `json:"category"`
	Content  string    `json:"content"`
	Creation time.Time `json:"creation"`
}
