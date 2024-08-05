default: run

run:
	@go run main.go

build: run
	@go build -o ./app
