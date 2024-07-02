postgres:
	docker run --name go-master-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:16.3-alpine

createdb:
	docker exec -it go-master-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it go-master-postgres dropdb simple_bank

migrateup:
	migrate -path simplebank/db/migration -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path simplebank/db/migration -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate
.PHONY: createdb dropdb postgres migrateup migratedown sqlc