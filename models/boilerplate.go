package models

import (
	"svc-boilerplate-golang/entity"
	"time"
)

type Boilerplate struct {
	entity.BoilerplateData
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
	entity.BoilerplateData
}

type BoilerplatePayloadUpdate struct {
	Param BoilerplateParamUpdate `json:"param" binding:"required"`
	Data  BoilerplateDataUpdate  `json:"data" binding:"required"`
}

type BoilerplateParamUpdate struct {
	entity.BoilerplateParameterUpdate
}

type BoilerplateDataUpdate struct {
	entity.BoilerplateData
}

type BoilerplatePayloadDetele struct {
	entity.BoilerplateParameterDelete
}
