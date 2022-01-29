package database

import (
	"strings"
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
}

func (cfg QueryConfig) QueryBuilder() {
	if cfg.Action == "select" {
		cfg.selectBuilder()
		cfg.whereBuilder(cfg.OnSelect.Where)
	} else if cfg.Action == "insert" {
		cfg.insertBuilder()
	} else if cfg.Action == "update" {
		cfg.updateBuilder()
		cfg.whereBuilder(cfg.OnUpdate.Where)
	} else if cfg.Action == "delete" {
		cfg.deleteBuilder()
		cfg.whereBuilder(cfg.OnDelete.Where)
	}
}

func (cfg QueryConfig) selectBuilder() {
	cfg.Result.Query += `SELECT `

	for _, x := range cfg.OnSelect.Column {
		cfg.Result.Query += x + ", "
	}

	cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-1] // TRIM THE LAST `,`

	cfg.Result.Query += `FROM ` + cfg.Table
}

func (cfg QueryConfig) insertBuilder() {
	cfg.Result.Query += `INSERT INTO ` + cfg.Table

	for _, x := range cfg.OnInsert.Data {
		count := len(x.([]interface{}))

		if count < 0 {
			count = 0
		}

		cfg.Result.Query += ` (` + cfg.getQuestionMark() + strings.Repeat(","+cfg.getQuestionMark(), count) + `),`

		cfg.Result.Value = append(cfg.Result.Value, x.([]interface{}))
	}

	cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-1] // TRIM THE LAST `,`
}

func (cfg QueryConfig) updateBuilder() {
	cfg.Result.Query += `UPDATE ` + cfg.Table + ` SET `

	for i, x := range cfg.OnUpdate.Data {
		cfg.Result.Query += i + ` = ` + cfg.getQuestionMark() + `,`
		cfg.Result.Value = append(cfg.Result.Value, x)
	}

	cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-1] // TRIM THE LAST `,`
}

func (cfg QueryConfig) deleteBuilder() {
	cfg.Result.Query += `SELECT FROM ` + cfg.Table
}

func (cfg QueryConfig) whereBuilder(param map[string]interface{}) {
	c := 0

	if len(param) > 0 {
		cfg.Result.Query += ` WHERE `
	}

	for i, x := range param {
		if i == "AND" {
			for g, v := range x.(map[string]interface{}) {
				if g == "IN" {
					for o, f := range v.(map[string]interface{}) {
						r := len(f.([]string)) - 1
						if r < 0 {
							r = 0
						}
						cfg.Result.Query += o + ` IN (` + cfg.getQuestionMark() + strings.Repeat(","+cfg.getQuestionMark(), r) + `) AND `
						for _, w := range f.([]string) {
							cfg.Result.Value = append(cfg.Result.Value, w)
						}
						c++
					}
				} else {
					cfg.Result.Query += g + ` = ` + cfg.getQuestionMark() + ` AND `
					cfg.Result.Value = append(cfg.Result.Value, v)
					c++
				}
			}
			cfg.Result.Query = cfg.Result.Query[0 : len(cfg.Result.Query)-4] // TRIM THE LAST `AND `
		}
	}
}

func (cfg QueryConfig) getQuestionMark() (questionMark string) {
	switch cfg.TechStack {
	case "oracle":
		questionMark = ":"
	case "mysql":
		questionMark = "?"
	}
	return
}
