FROM golang:1.20-alpine as builder
WORKDIR /build

COPY . .
EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hanlder ./cmd/main.go

FROM alpine as production
WORKDIR /app
COPY --from=builder /build/ ./
CMD ["/app/handler"]