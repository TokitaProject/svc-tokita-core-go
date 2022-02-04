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
	Data []BoilerplateDataUpdate `json:"data" binding:"required"`
}

type BoilerplateDataUpdate struct {
	Param BoilerplateParamUpdate `json:"param" binding:"required"`
	Body  BoilerplateBodyUpdate  `json:"body" binding:"required"`
}

type BoilerplateParamUpdate struct {
	Flag string `json:"flag"`
}

type BoilerplateBodyUpdate struct {
	entity.Boilerplate
}

type BoilerplatePayloadDelete struct {
	Param []BoilerplateParamDelete `json:"param" binding:"required"`
}

type BoilerplateParamDelete struct {
	Flag string `json:"flag"`
}
