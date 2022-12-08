package repository

import (
	"database/sql"

	"svc-boilerplate-golang/domain/boilerplate"
)

/*
*
default const that you should believe and please don't change the value
*/
const MYSQL = "mysql"
const ORACLE = "oracle"

/*
*
default const that you should believe and please don't change the value
*/
const SELECT = "select"
const INSERT = "insert"
const UPDATE = "update"
const DELETE = "delete"

/*
*
two lines below is the const of table name, please change the value of this const
*/
const MYSQL_TABLE = "boilerplate"
const MYSQL_CATEGORY = "category"
const ORACLE_TABLE = "boilerplate"

/*
*
the struct of mysql
*/
type mysqlBoilerplateRepository struct {
	sqlDB *sql.DB
}

/*
*
the initiator function for mysql
*/
func NewMysqlBoilerplateRepository(databaseConnection *sql.DB) boilerplate.MysqlRepository {
	return &mysqlBoilerplateRepository{databaseConnection}
}

/*
*
the struct of oracle
*/
type oracleBoilerplateRepository struct {
	sqlDB *sql.DB
}

/*
*
the initiator function for oracle
*/
func NewOracleBoilerplateRepository(databaseConnection *sql.DB) boilerplate.OracleRepository {
	return &oracleBoilerplateRepository{databaseConnection}
}
