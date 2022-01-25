package entity

import "time"

type UUID struct {
	ID   uint64 `json:"-"`
	UUID string `json:"uuid"`
}

type Time struct {
	UserInput     string     `json:"user_input"`
	TanggalInput  time.Time  `json:"tgl_input"`
	UserUpdate    *string    `json:"user_update"`
	TanggalUpdate *time.Time `json:"tgl_update"`
}
