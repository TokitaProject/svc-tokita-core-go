package entity

type BoilerplateKey struct {
	ForeignID uint64 `json:"-"`
}

type Boilerplate struct {
	BoilerplateKey
	Categotry  string `json:"category_id"`
	Name       string `json:"name"`
	LastUpdate string `json:"last_update"`
}
