




FROM golang:1.19 as builder
COPY . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o start .

FROM scratch as production
WORKDIR /opt/app
COPY --from=builder /build/ ./
CMD ["/opt/app/start"]

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o handler ./cmd/main.go