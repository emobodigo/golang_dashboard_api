migrate_up:
	migrate -database "mysql://root@tcp(localhost:3306)/dashboard_api" -path db/migrations up

migrate_down:
	migrate -database "mysql://root@tcp(localhost:3306)/dashboard_api" -path db/migrations down

migrate_force:
	migrate -database "mysql://root@tcp(localhost:3306)/dashboard_api" -path db/migrations force

run:
	go build && ./golang_dashboard_api

migrate_test:
	migrate -database "mysql://root@tcp(localhost:3306)/dashboard_api_test" -path db/migrations up

migrate_ci:
	migrate -database "mysql://root:root@tcp(localhost:3306)/dashboard_api_test" -path db/migrations up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrate_up migrate_down migrate_force run migrate_test sqlc migrate_ci test