export DB_DRIVERNAME="mysql"
export DB_USERNAME="root"
export DB_PASSWORD="shindi"
export DB_HOST="localhost"
export DB_PORT="3306"
export DB_NAME="svc_boilerlate_golang"

export PORT="80"

export GIN_MODE=debug

nodemon --exec go run main.go --signal SIGTERM