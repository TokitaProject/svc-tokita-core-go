package valueobject

import (
	"svc-boilerplate-golang/entity"
)

type Boilerplate struct {
	entity.Boilerplate
	entity.StandardKey
	entity.Time
}

type BoilerplatePayloadInsert struct {
	Data []BoilerplateDataInsert `json:"data" binding:"required"`
}

type BoilerplateDataInsert struct {
	entity.Boilerplate
}

type BoilerplatePayloadUpdate struct {
	Param BoilerplateParamUpdate `json:"param" binding:"required"`
	Data  BoilerplateDataUpdate  `json:"data" binding:"required"`
}

type BoilerplateParamUpdate struct {
	Flag string `json:"flag"`
}

type BoilerplateDataUpdate struct {
	entity.Boilerplate
}

type BoilerplatePayloadDetele struct {
	Flag string `json:"flag" binding:"required"`
}
