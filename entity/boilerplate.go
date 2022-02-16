package entity

type BoilerplateKey struct {
	ForeignID uint64 `json:"foreign_id"`
}

type Boilerplate struct {
	BoilerplateKey
	Flag   string `json:"flag"`
	Column string `json:"column"`
}
