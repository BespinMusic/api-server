export DBNAME=lintest 
docker-compose up -d
sleep 5
go run migrations/migrate.go $DBNAME
TEST_DB_USERNAME=postgres TEST_DB_PASSWORD=test TEST_DB_NAME=$DBNAME go test -v
docker-compose down