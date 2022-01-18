package models

type Boilerplate struct {
	ID   string `json:"-"`
	UUID string `json:"uuid"`
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
