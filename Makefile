postgres:
	docker run --name go-backend-practice -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:alpine

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable" -verbose down

createdb:
	docker exec -it go-backend-practice createdb --username postgres --owner postgres bank
	
dropdb:
	docker exec -it go-backend-practice dropdb bank                                 

.PHONY: postgres migrateup migratedown