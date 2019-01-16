.PHONY: proto

proto:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/gotut/vessel-service \
		proto/vessel/vessel.proto

build:
	docker build -t vessel-service .

run:
	docker run --rm -p 8001:8000 \
		-e MICRO_SERVER_ADDRESS=:8000 \
		-e MICRO_REGISTRY=mdns vessel-service