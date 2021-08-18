# FROM golang:1.16-alpine3.13
FROM golang:1.16

WORKDIR /go/src/app
COPY . .

# RUN go get -d -v ./...
# RUN cp worker-launch /go/bin/worker-launch
# RUN chmod +x /go/bin/worker-launch

RUN mkdir /go/bin/config-server
RUN go get -d -v ./...
RUN cp config-server /go/bin/config-server/config-server
# RUN go install -v ./...

# CMD ["/go/bin/worker-launch"]
CMD ["./config-server", "-deploy"]