package model

import "time"

type Todo struct {
	ID                uint        `gorm:"primary_key" json:"id" `
	Title             string      `json:"title" binding:"required,max=127"`
	Content           string      `json:"content"`
	Status            string      `json:"status"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
}