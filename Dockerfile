FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o go

ENTRYPOINT ["/app/go"]