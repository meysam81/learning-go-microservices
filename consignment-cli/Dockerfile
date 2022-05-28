FROM debian:11

RUN mkdir -p /app
WORKDIR /app

ADD consignment.json /app/consignment.json
ADD consignment-cli.out /app/consignment-cli

CMD ["./consignment-cli"]
