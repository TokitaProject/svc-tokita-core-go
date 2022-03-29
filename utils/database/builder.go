package database

import (
	"strconv"
)

type OnSelect struct {
	Column []string
	Where  map[string]interface{}
}

type OnInsert struct {
	Column []string
	Data   []interface{}
}

type OnUpdate struct {
	Where map[string]interface{}
	Data  map[string]interface{}
}

type OnDelete struct {
	Where map[string]interface{}
}

type TableInfo struct {
	TechStack string
	Table     string
	Action    string
}

type Result struct {
	Query string
	Value []interface{}
}

type QueryConfig struct {
	TableInfo
	OnSelect
	OnInsert
	OnUpdate
	OnDelete
	Result
	counter int
}

func (cfg *QueryConfig) QueryBuilder() {
	cfg.counter = 0
	if cfg.Action == "select" {
		cfg.selectBuilder()
		cfg.whereBuilder(cfg.OnSelect.Where)
	} else if cfg.Action == "select-distinct" {
		cfg.selectDistinctBuilder()
		cfg.whereBuilder(cfg.OnSelect.Where)
	} else if cfg.Action == "insert" {
		if cfg.TechStack == "oracle" && len(cfg.OnInsert.Data) > 1 {
			cfg.insertOracleBatchBuilder()
		} else {
			cfg.insertBuilder()
		}
	} else if cfg.Action == "update" {
		cfg.updateBuilder()
		cfg.whereBuilder(cfg.OnUpdate.Where)
	} else if cfg.Action == "delete" {
		cfg.deleteBuilder()
		cfg.whereBuilder(cfg.OnDelete.Where)
	}
}

func (cfg *QueryConfig) selectBuilder() {
	cfg.Result.Query += `SELECT `

	for _, x := range cfg.OnSelect.Column {
		cfg.Result.Query += x + ", "
	}

	cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-2]

	cfg.Result.Query += ` FROM ` + cfg.Table
}

func (cfg *QueryConfig) selectDistinctBuilder() {
	cfg.Result.Query += `SELECT DISTINCT `

	for _, x := range cfg.OnSelect.Column {
		cfg.Result.Query += x + ", "
	}

	cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-2]

	cfg.Result.Query += ` FROM ` + cfg.Table
}

func (cfg *QueryConfig) insertBuilder() {
	cfg.Result.Query += `INSERT INTO ` + cfg.Table + ` (`

	for _, x := range cfg.OnInsert.Column {
		cfg.Result.Query += x + `, `
	}

	cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-2]
	cfg.Result.Query += `) VALUES `

	for _, x := range cfg.OnInsert.Data {
		count := len(x.([]interface{}))

		if count < 0 {
			count = 0
		}

		cfg.Result.Query += `(`
		for i := 0; i < count; i++ {
			cfg.Result.Query += cfg.getQuestionMark() + `, `
		}
		cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-2]
		cfg.Result.Query += `),`

		cfg.Result.Value = append(cfg.Result.Value, x.([]interface{})...)
	}

	cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-1]
}

func (cfg *QueryConfig) insertOracleBatchBuilder() {
	cfg.Result.Query += `INSERT ALL`

	for _, x := range cfg.OnInsert.Data {
		count := len(x.([]interface{}))

		if count < 0 {
			count = 0
		}

		cfg.Result.Query += ` INTO ` + cfg.Table + `(`

		for _, x := range cfg.OnInsert.Column {
			cfg.Result.Query += x + `, `
		}

		cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-2]
		cfg.Result.Query += `) VALUES `

		cfg.Result.Query += `(`
		for i := 0; i < count; i++ {
			cfg.Result.Query += cfg.getQuestionMark() + `, `
		}
		cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-2]
		cfg.Result.Query += `) `

		cfg.Result.Value = append(cfg.Result.Value, x.([]interface{})...)
	}

	cfg.Result.Query += `SELECT * FROM dual`
}

func (cfg *QueryConfig) updateBuilder() {
	cfg.Result.Query += `UPDATE ` + cfg.Table + ` SET `

	for i, x := range cfg.OnUpdate.Data {
		cfg.Result.Query += i + ` = ` + cfg.getQuestionMark() + `, `
		cfg.Result.Value = append(cfg.Result.Value, x)
	}

	cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-2]
}

func (cfg *QueryConfig) deleteBuilder() {
	cfg.Result.Query += `DELETE FROM ` + cfg.Table
}

func (cfg *QueryConfig) whereBuilder(param map[string]interface{}) {
	found := false

	cfg.Result.Query += ` WHERE `

	for i, x := range param {
		if i == "AND" {
			for g, v := range x.(map[string]interface{}) {
				if g == "IN" {
					for o, f := range v.(map[string]interface{}) {
						r := len(f.([]string))
						if r < 1 {
							continue
						}
						cfg.Result.Query += o + ` IN (`
						for i := 0; i < r; i++ {
							if f.([]string)[i] == "" {
								continue
							}
							cfg.Result.Query += cfg.getQuestionMark() + `, `
						}
						cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-2]
						cfg.Result.Query += `) AND `
						for _, w := range f.([]string) {
							if w == "" {
								continue
							}
							cfg.Result.Value = append(cfg.Result.Value, w)
						}
						found = true
					}
				} else if g == "NOT" {
					for o, f := range v.(map[string]interface{}) {
						if f == "" {
							continue
						}
						if f == nil {
							cfg.Result.Query += o + ` IS NOT NULL AND `
							found = true
						} else {
							cfg.Result.Query += `NOT ` + o + ` = ` + cfg.getQuestionMark() + ` AND `
							cfg.Result.Value = append(cfg.Result.Value, f)
							found = true
						}
					}
				} else {
					if v == "" {
						continue
					}
					if v == nil {
						cfg.Result.Query += g + ` IS NULL AND `
						found = true
					} else {
						cfg.Result.Query += g + ` = ` + cfg.getQuestionMark() + ` AND `
						cfg.Result.Value = append(cfg.Result.Value, v)
						found = true
					}
				}
			}
			if found {
				cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-4]
			}
		}
	}

	if !found {
		cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-7]
	}
}

func (cfg *QueryConfig) getQuestionMark() (questionMark string) {
	switch cfg.TechStack {
	case "oracle":
		questionMark = ":x" + strconv.Itoa(cfg.counter)
		cfg.counter++
	case "mysql":
		questionMark = "?"
	}
	return
}
