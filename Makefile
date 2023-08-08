migrate_up:
	migrate -database "mysql://root@tcp(localhost:3306)/dashboard_api" -path db/migrations up

migrate_down:
	migrate -database "mysql://root@tcp(localhost:3306)/dashboard_api" -path db/migrations down

migrate_force:
	migrate -database "mysql://root@tcp(localhost:3306)/dashboard_api" -path db/migrations force

run:
	go build && ./golang_dashboard_api

migrate_up_test:
	migrate -database "mysql://root@tcp(localhost:3306)/dashboard_api_test" -path db/migrations up

migrate_down_test:
	migrate -database "mysql://root@tcp(localhost:3306)/dashboard_api_test" -path db/migrations down

migrate_force_test:
	migrate -database "mysql://root@tcp(localhost:3306)/dashboard_api_test" -path db/migrations force