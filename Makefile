migrate:
	migrate -path ./db/schema -database 'postgres://postgres:mysecretpassword@0.0.0.0:5432/postgres?sslmode=disable' up

run:
	docker-compose up --build 