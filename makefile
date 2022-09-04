postgres :
	docker run --name sls -p 5432:5432 --network slsnet -e POSTGRES_PASSWORD=mysecretpassword -d postgres:14.5-alpine

startdb:
	docker exec -it sls psql -U postgres

createdb:
	docker exec -it sls createdb slsstore -U postgres

dropdb:
	docker exec -it sls dropdb slsstore -U postgres

migratedb:
	docker exec

test:
	go test -v -cover ./...

build:
	go build



.PHONY: postgres startdb