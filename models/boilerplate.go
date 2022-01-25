package models

import "time"

type Boilerplate struct {
	ID            uint64     `json:"-"`
	UserInput     string     `json:"user_input"`
	TanggalInput  time.Time  `json:"tgl_input"`
	UserUpdate    *string    `json:"user_update"`
	TanggalUpdate *time.Time `json:"tgl_update"`
	UUID          string     `json:"uuid"`
}

type BoilerplatePayloadInsert struct {
	Data []BoilerplateDataInsert `json:"data" binding:"required"`
}

type BoilerplateDataInsert struct {
	Column string `json:"column"`
}

type BoilerplatePayloadUpdate struct {
	Param BoilerplateParamUpdate `json:"param" binding:"required"`
	Data  BoilerplateDataUpdate  `json:"data" binding:"required"`
}

type BoilerplateParamUpdate struct {
	Flag int `json:"flag"`
}

type BoilerplateDataUpdate struct {
	Column string `json:"column"`
}

type BoilerplatePayloadDetele struct {
	Flag int `json:"flag" binding:"required"`
}
