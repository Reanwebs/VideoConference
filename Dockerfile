FROM golang:1.21.0

WORKDIR /app

ADD . /app

WORKDIR /app/cmd

COPY .env /app/cmd/.env

RUN go mod download

RUN go test -v ./...

RUN go build -o main .

CMD ["/app/cmd/main"]