FROM golang:1.12
ENV GO111MODULE=on
WORKDIR /request-log

COPY go.mod .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 8080

ENTRYPOINT ["/request-log/request-log"]