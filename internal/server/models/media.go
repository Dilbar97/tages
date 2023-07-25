package models

import (
	"time"
)

type Media struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ListResponse struct {
	Success bool    `json:"success"`
	Data    []Media `json:"data,omitempty"`
	Error   string  `json:"error,omitempty"`
}

type GetMediaResponse struct {
	Success bool   `json:"success"`
	Data    Media  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
