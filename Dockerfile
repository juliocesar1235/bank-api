FROM golang:1.22.0-alpine as builder
COPY go.mod go.sum /go/src/github.com/juliocesar1235/bank-api/
WORKDIR /go/src/github.com/juliocesar1235/bank-api
RUN go mod download
COPY . /go/src/github.com/juliocesar1235/bank-api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/bank-api github.com/juliocesar1235/bank-api

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/juliocesar1235/bank-api/build/bank-api /usr/bin/bank-api
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/bank-api"]