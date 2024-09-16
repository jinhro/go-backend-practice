migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable" -verbose down
	
.PHONY: migrateup migratedown