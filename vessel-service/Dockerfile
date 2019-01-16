FROM golang:latest as builder

WORKDIR /go/src/gotut/vessel-service

COPY . .

RUN go get -v -u github.com/golang/dep/cmd/dep
RUN go get -v -u github.com/golang/protobuf/protoc-gen-go
RUN dep init -v && dep ensure -v
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o ${PWD##*/}.out   .

FROM alpine:latest

RUN apk --co-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/gotut/vessel-service/vessel-service.out .

CMD ["./vessel-service"]
