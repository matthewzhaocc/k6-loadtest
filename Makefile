buildserver:
	go build -o server .

startserver: buildserver
	GIN_MODE=release ./server

image:
	docker build . -t matthewzhaocc/k6-server:latest