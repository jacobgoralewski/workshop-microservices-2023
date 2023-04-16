
docker-build-dev:
	sudo docker build -t users .

proto:
	@protoc -I users --go_out=plugins=grpc:users users.proto
