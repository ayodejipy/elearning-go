run:
	go run cmd/app/main.go

build:
	go build -o bin/elearning cmd/app/main.go && bin/elearning