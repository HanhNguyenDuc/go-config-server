FROM golang:1.16-alpine3.13

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM golang:1.16-alpine3.13

WORKDIR /app

COPY --from=builder /app/config-server /usr/bin/

ENTRYPOINT ["config-server"]
