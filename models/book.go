package models

import "time"

type Book struct {
	Name       string    `json:"name"`
	Id         string    `json:"id"`
	Category   string    `json:"category"`
	Author     string    `json:"author"`
	Translator string    `json:"translator"`
	Owner      string    `json:"owner"`
	AddTime    time.Time `json:"add_time"`
}
