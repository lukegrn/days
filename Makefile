serve:
	go run ./cmd/serve/main.go

build:
	go build -o ./bin ./cmd/serve
	go build -o ./bin ./cmd/hashgen 
