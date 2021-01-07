FROM golang:1.10.8-alpine3.9

WORKDIR /go/src/app

COPY main.go .

RUN go build -o main .

EXPOSE 8888

CMD ["./main"]