## run: runs the api server
run-be::
	go run cmd/api/*.go

## db/psql: connects to the database using psql
db/psql:
	docker exec -it pet_shelter_db psql -U pet_admin -d shelter_db

## tidy: formats code and cleans dependencies
tidy:
	go fmt ./...
	go mod tidy