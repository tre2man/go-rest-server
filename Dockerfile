FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o main .

ENV GIN_MODE=release

EXPOSE 8080

CMD ["/app/main"] 

