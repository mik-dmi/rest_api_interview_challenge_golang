FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/api_golang cmd/main.go

EXPOSE 8080

CMD ["/go/bin/api_golang"]
