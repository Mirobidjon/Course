createdb:
    docker run --name=course-db -e POSTGRES_PASSWORD='secret' -p 5436:5432 -d postgres

migraterup:
    migrate -path ./schema database 'postgres://postgres:secret@localhost:5436/postgres?sslmode=disable' up

migratedown:
    migrate -path ./schema database 'postgres://postgres:secret@localhost:5436/postgres?sslmode=disable' down

dropdb:
    docker exec -it postrges dropdb course-db

run:
    go run cmd/main.go