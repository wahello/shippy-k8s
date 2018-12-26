FROM golang:1.11.4-alpine3.8 as builder
RUN apk add --update --no-cache ca-certificates git

RUN mkdir /server
WORKDIR /server

COPY go.* ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/server -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM scratch as release
COPY --from=builder /go/bin/server /go/bin/server
ENTRYPOINT ["/go/bin/server"]
