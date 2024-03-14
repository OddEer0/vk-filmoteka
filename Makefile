swagger:
	rm -rf ./docs
	swag init -g cmd/main/main.go

run:
	go run ./cmd/main/main.go

clean:
	rm -rf ./docs