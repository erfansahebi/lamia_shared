proto:
	protoc services/**/proto/*.proto --go_out=plugins=grpc:.

server:
	go run main.go