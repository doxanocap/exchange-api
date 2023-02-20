# Build the Go API
FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go mod download

RUN cd .\cmd\ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../auth.exe

EXPOSE 8000

CMD ["/auth"]