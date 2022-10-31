cli:
	go build -mod vendor -o bin/post cmd/post/main.go
	go build -mod vendor -o bin/api cmd/api/main.go
