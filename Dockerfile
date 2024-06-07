FROM golang:1.22-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build

EXPOSE 5555

CMD ["./main"]
