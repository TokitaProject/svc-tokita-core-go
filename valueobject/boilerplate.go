package valueobject

import "svc-boilerplate-golang/entity"

type Boilerplate struct {
	entity.Boilerplate
	entity.StandardKey
	entity.Time
}

type BoilerplatePayloadInsert struct {
	Data []Boilerplate `json:"data" binding:"required"`
}

type BoilerplatePayloadUpdate struct {
	Data []BoilerplateDataUpdate `json:"data" binding:"required"`
}

type BoilerplateDataUpdate struct {
	Param Boilerplate `json:"param" binding:"required"`
	Body  Boilerplate `json:"body" binding:"required"`
}

type BoilerplatePayloadDelete struct {
	Param []Boilerplate `json:"param" binding:"required"`
}
