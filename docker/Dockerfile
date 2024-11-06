FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main cmd/main.go

EXPOSE 8080 50051 8081

CMD [ "/app/main" ]
