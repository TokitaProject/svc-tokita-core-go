package valueobject

import "svc-boilerplate-golang/entity"

type Boilerplate struct {
	entity.Boilerplate
	entity.StandardKey
	entity.Time
}

type BoilerplatePayloadInsert struct {
	Data []Boilerplate `json:"data" binding:"required"`
	User string
}

type BoilerplatePayloadUpdate struct {
	Data []BoilerplateDataUpdate `json:"data" binding:"required"`
	User string
}

type BoilerplateDataUpdate struct {
	Param Boilerplate `json:"param" binding:"required"`
	Body  Boilerplate `json:"body" binding:"required"`
}

type BoilerplatePayloadDelete struct {
	Param []Boilerplate `json:"param" binding:"required"`
}

type Nilai struct {
	entity.Nilai
	entity.StandardKey
	entity.Time
}

type NilaiPayloadInsert struct {
	Data []Nilai `json:"data" binding:"required"`
	User string
}

type NilaiPayloadUpdate struct {
	Data []NilaiDataUpdate `json:"data" binding:"required"`
	User string
}

type NilaiDataUpdate struct {
	Param Nilai `json:"param" binding:"required"`
	Body  Nilai `json:"body" binding:"required"`
}

type NilaiPayloadDelete struct {
	Param []Nilai `json:"param" binding:"required"`
}
