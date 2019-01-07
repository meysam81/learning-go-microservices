proto:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/gotut/consignment-service \
	  proto/consignment/consignment.proto

build:
	go build -o consignment-service.out main.go
	docker build -t consignment-service .

run:
	docker run -dp 8000:8000 consignment-service
