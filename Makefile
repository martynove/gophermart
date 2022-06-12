migrate_up:
	migrate -path ./migrations/schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./migrations/schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down

run_go_debug:
	go run cmd/gophermart/main.go -a localhost:8080 -d postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable -r localhost:8081 --debug

run_go:
	go run cmd/gophermart/main.go -a localhost:8080 -d postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable -r localhost:8081

docker_psql_start:
	docker run --name=gophermart-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
docker_psql_stop:
	docker stop gophermart-db
