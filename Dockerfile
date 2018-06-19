# build stage
FROM golang:latest AS build-env
WORKDIR /go/src/github.com/bespinmusic/api-server
ADD . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goapp .
WORKDIR /go/src/github.com/bespinmusic/api-server/migrations
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migrate migrate.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/bespinmusic/api-server/goapp .
COPY --from=build-env /go/src/github.com/bespinmusic/api-server/migrations ./migrations
CMD ["./goapp"]