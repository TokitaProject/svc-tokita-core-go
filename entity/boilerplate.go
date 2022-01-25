package entity

type BoilerplateData struct {
	Column string `json:"column"`
}

type BoilerplateParameterUpdate struct {
	Flag string `json:"flag"`
}

type BoilerplateParameterDelete struct {
	Flag string `json:"flag" binding:"required"`
}
