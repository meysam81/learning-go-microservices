FROM golang:latest as builder

WORKDIR /go/src/gotut/consignment-service
COPY . .

RUN go get -v -u github.com/golang/dep/cmd/dep
RUN dep init -v && dep ensure -v
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o consignment-service.out .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/gotut/consignment-service .
CMD ["./consignment-service.out"]
