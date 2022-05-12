package entity

type BoilerplateKey struct {
	ForeignID uint64 `json:"-"`
}

type Boilerplate struct {
	BoilerplateKey
	Flag   string `json:"flag,omitempty"`
	Column string `json:"column,omitempty"`
}
