.PHONY: proto

proto:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/gotut/consignment-service \
		proto/consignment/consignment.proto

build:
	go build -o consignment-service.out
	docker build -t consignment-service .

run:
	docker network ls | grep "consignment" | [ `wc -l` -ne 1 ] && \
		docker network create consignment || echo '`consignment` network already exists'
	docker run --net consignment -dp 8000:8000 \
		-e MICRO_SERVER_ADDRESS=:8000 \
		-e MICRO_REGISTRY=mdns consignment-service
