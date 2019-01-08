FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/meysam81/learning-go-microservices/consignment-service
COPY . .

RUN go get -v -u github.com/golang/dep/cmd/dep
RUN dep init -v && dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/gotut/consignment-service .
CMD ["./consignment-service"]
