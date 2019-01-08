proto:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/gotut/consignment-service \
	  proto/consignment/consignment.proto

build:
	go build -o consignment-service.out
	docker build -t consignment-service .

run:
	docker run -dp 8000:8000 -e MICRO_SERVER_ADDRESS=:8000 \
		-e MICRO_REGISTRY=mdns consignment-service
