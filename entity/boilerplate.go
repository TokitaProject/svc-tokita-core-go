package entity

type BoilerplateKey struct {
	ForeignID uint64 `json:"-"`
}

type Boilerplate struct {
	BoilerplateKey
	Email string `json:"email"`
	Kelas string `json:"kelas"`
	Nis   string `json:"nis"`
	Nama  string `json:"nama"`
}

type Nilai struct {
	Nis        string `json:"nis"`
	BhsInd     string `json:"bhs_ind"`
	BhsInggris string `json:"bhs_inggris"`
	Mtk        string `json:"mtk"`
}
