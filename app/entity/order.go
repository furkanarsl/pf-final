package entity

import "time"

type Order struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	OrderedAt time.Time `json:"ordered_at"`
	Total     float64   `json:"total"`
}
