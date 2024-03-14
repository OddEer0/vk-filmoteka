swagger:
	rm -rf ./docs
	swag init -g cmd/main/main.go

run:
	CONFIG_PATH=./config/dev.yaml go run ./cmd/main/main.go

clean:
	rm -rf ./docs