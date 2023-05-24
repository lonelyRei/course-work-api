# To run backend:
1. docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
2. migrate -path ./scema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
3. go run .\cmd\main.go 
