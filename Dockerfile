FROM golang:1.17.2-bullseye

ENV GO111MODULE=auto

WORKDIR /go/api

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build main.go
ENV GIN_MODE=release
EXPOSE 80

# [TODO] use alpine

CMD ["/go/api/main"]
