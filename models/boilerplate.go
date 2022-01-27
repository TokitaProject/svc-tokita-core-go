package models

import (
	"svc-boilerplate-golang/entity"
)

type Boilerplate struct {
	entity.BoilerplateData
	entity.StandardKey
	entity.Time
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
