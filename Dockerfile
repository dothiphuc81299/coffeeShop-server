FROM golang:alpine

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build cmd/admin/main.go

EXPOSE 5000

CMD ["./main"]

 