package entity

import "time"

type StandardKey struct {
	ID   uint64 `json:"-"`
	UUID string `json:"uuid,omitempty"`
}

type Time struct {
	UserInput     string     `json:"user_input,omitempty"`
	TanggalInput  *time.Time `json:"tgl_input,omitempty"`
	UserUpdate    string     `json:"user_update,omitempty"`
	TanggalUpdate *time.Time `json:"tgl_update,omitempty"`
}
