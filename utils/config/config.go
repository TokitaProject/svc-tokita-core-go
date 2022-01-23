package config

var Config = map[string]string{
	"sql.not.found": "sql: no rows in result set",
}

func Get(name string) string {
	return Config[name]
}
