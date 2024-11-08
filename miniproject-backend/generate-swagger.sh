rm -rf ./docs

swag init -g cmd/api/main.go --output docs

go fmt ./docs/...
