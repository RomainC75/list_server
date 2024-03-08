migrateup:
	migrate --path db/migration --database "postgresql://name:pass@localhost:5432/test?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migration --database "postgresql://name:pass@localhost:5432/test?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrateup migratedown sqlc test