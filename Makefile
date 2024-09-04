VERSION=0.0

all:
	go build
	GOOS=windows GOARCH=amd64 go build
