FROM golang:1.22-alpine

WORKDIR /app

COPY * ./

RUN go mod download main 

RUN CGO_ENABLED=0 GOOS=linux go build main.go

EXPOSE 5555

CMD ["./main.go"]

