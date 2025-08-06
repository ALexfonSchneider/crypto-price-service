
.PHONY: sqlc-install
sqlc-install:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

.PHONY: sqlc
sqlc:
	sqlc generate

swagger-install:
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: swagger
swagger:
	swag init \
	  --dir . \
	  --output docs \
	  --parseInternal \
	  --parseDependency \
	  --generalInfo cmd/service/main.go

.PHONY: run
run:
	go run cmd/service/main.go

migrate:
	go run cmd/migrate/main.go