run:
	@go run ./cmd/main.go
build:
	@go build -o ./bin/gme_server_app ./cmd/main.go
daemon:
	@nohup ./bin/gme_server_app > serverlogs.txt
