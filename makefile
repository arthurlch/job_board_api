postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb: 
	sleep 5
	docker exec -it postgres12 createdb --username=root --owner=root jobb_dev

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/jobb_dev?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/jobb_dev?sslmode=disable" -verbose down

dropdb: 
	docker exec -it postgres12 dropdb jobb_dev

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

.PHONY: postgres createdb dropdb migrateup migratedown