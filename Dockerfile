# build stage
FROM golang:latest AS build-env
WORKDIR /go/src/github.com/bespinmusic/api-server
ADD . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goapp .

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/bespinmusic/api-server/goapp .
CMD ["./goapp"]