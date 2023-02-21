# Build the Go API
FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hanlder.exe ./cmd/main.go

EXPOSE 8080

CMD ["./handler.exe"]

