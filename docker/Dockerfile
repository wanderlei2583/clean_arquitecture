FROM golang:1.23-alpine

WORKDIR /app

# Instalar dependências necessárias
RUN apk add --no-cache protobuf-dev protoc git gcc musl-dev

# Instalar ferramentas do protobuf
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Copiar apenas o go.mod e go.sum primeiro
COPY go.mod go.sum ./
RUN go mod download

# Copiar o resto dos arquivos
COPY . .

# Criar diretório pb se não existir
RUN mkdir -p api/grpc/pb

# Gerar arquivos protobuf
RUN protoc \
    --proto_path=api/grpc/proto \
    --go_out=api/grpc/pb --go_opt=paths=source_relative \
    --go-grpc_out=api/grpc/pb --go-grpc_opt=paths=source_relative \
    api/grpc/proto/pedido.proto

# Compilar
RUN go build -o main cmd/main.go

EXPOSE 8080 50051 8081

CMD ["./main"]
