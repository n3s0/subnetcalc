VERSION=0.0

all:
	GOOS=linux GOARCH=amd64 go build
	GOOS=windows GOARCH=amd64 go build
