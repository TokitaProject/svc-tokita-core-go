package config

var Config = map[string]string{
	"sql.not.found":   "sql: no rows in result set",
	"error.bind.json": "parameter tidak tepat",
}

func Get(name string) string {
	return Config[name]
}
