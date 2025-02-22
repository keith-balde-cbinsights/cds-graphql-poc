
generate:
	go generate ./...

generate-gqlgen:
	go run github.com/99designs/gqlgen generate

run:
	go run server.go

