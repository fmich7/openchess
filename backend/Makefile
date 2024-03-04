build:
	go build -o bin/app ./cmd/api/main.go 

run: build
	./bin/app

createDB:
	@docker run --name postgres -e POSTGRES_PASSWORD=admin -e POSTGRES_USER=postgres -p 5432:5432 -d postgres

test:
	go test -v ./... -count=1