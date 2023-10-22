package model

import "time"

type Base struct {
	ID        string    `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at valid:"-"`
	UpdateAt  time.Time `json:"updated_at valid:"-"`
}
