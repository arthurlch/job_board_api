postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb: 
	docker exec -it postgres12 createdb --username=root --owner=root jobb_dev

dropdb: 
	docker exec -it postgres12 dropdb jobb_dev

.PHONY: postgres createdb dropdb