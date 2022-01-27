export DB_DRIVERNAME="mysql"
export DB_USERNAME="root"
export DB_PASSWORD="shindi"
export DB_HOST="localhost"
export DB_PORT="3306"
export DB_NAME="svc_boilerlate_golang"

export ORACLE_DB_DRIVERNAME="godror"
export ORACLE_DB_USERNAME="keubank"
export ORACLE_DB_PASSWORD="testing#"
export ORACLE_DB_HOST="10.30.21.17"
export ORACLE_DB_PORT="1521"
export ORACLE_DB_SERVICE_NAME="transdb1"

export PORT="80"

export GIN_MODE=debug

nodemon --exec go run main.go --signal SIGTERM