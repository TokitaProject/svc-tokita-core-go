# export DB_DRIVERNAME="mysql"
# export DB_USERNAME="root"
# export DB_PASSWORD="river"
# export DB_HOST="localhost"
# export DB_PORT="3306"
# export DB_NAME="svc_boilerlate_golang"

# export ORACLE_DB_DRIVERNAME="godror"
# export ORACLE_DB_USERNAME="keubank"
# export ORACLE_DB_PASSWORD="testing#"
# export ORACLE_DB_HOST="10.30.21.17"
# export ORACLE_DB_PORT="1521"
# export ORACLE_DB_SERVICE_NAME="transdb1"

export PORT="80"
export ENV="local"
export GIN_MODE=debug

# Bearer digunakan untuk melakukan request public api. Variabel env ini digunakan hanya untuk keperluan local development. Tidak ada dokumentasi, silakan bertanya.
export BEARER="bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJzdmMtbG9naW4tdWlpIiwic3ViIjoiZWxiby5zaGluZGktUEFva1dvS2lYbTY5UnBrZTVOWG04VGxLa3h2ckcxSDYiLCJpYXQiOjE2NDQyMTk5NTYsImV4cCI6MTY0NTQyOTU1Nn0.reAb-WHHimHD-CMd9otYVC4iIMFaJJ1ex-SGASYNzQQ"

nodemon --exec go run main.go --signal SIGTERM