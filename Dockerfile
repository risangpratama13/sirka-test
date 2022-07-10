FROM golang:1.18.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY app.env ./

RUN go mod download

COPY . .

RUN go build -o ./output/sirka-test

EXPOSE 8080

CMD ["./output/sirka-test"]