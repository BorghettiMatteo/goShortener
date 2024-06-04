FROM docker.io/golang:1.22

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build 

EXPOSE 8080

CMD ["./main"]