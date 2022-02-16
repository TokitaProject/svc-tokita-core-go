package repository

import (
	"database/sql"

	"svc-boilerplate-golang/domain/boilerplate"
)

const MYSQL = "mysql"
const ORACLE = "oracle"

const SELECT = "select"
const INSERT = "insert"
const UPDATE = "update"
const DELETE = "delete"

const MYSQL_TABLE = "boilerplate"
const ORACLE_TABLE = "boilerplate"

type mysqlBoilerplateRepository struct {
	sqlDB *sql.DB
}

func NewMysqlBoilerplateRepository(databaseConnection *sql.DB) boilerplate.MysqlRepository {
	return &mysqlBoilerplateRepository{databaseConnection}
}

type oracleBoilerplateRepository struct {
	sqlDB *sql.DB
}

func NewOracleBoilerplateRepository(databaseConnection *sql.DB) boilerplate.OracleRepository {
	return &oracleBoilerplateRepository{databaseConnection}
}
