build:
	go build -o consignment-cli.out
	docker build -t consignment-cli .

run:
	docker run -e MICRO_REGISTRY=mdns consignment-cli
