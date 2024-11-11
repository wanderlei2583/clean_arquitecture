package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/wanderlei2583/clean_arquitecture/api/grpc/pb"
	grpcserver "github.com/wanderlei2583/clean_arquitecture/api/grpc/server"
	"github.com/wanderlei2583/clean_arquitecture/internal/repository"
	"github.com/wanderlei2583/clean_arquitecture/internal/usecase"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=db_pedidos port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	pedidoRepo := repository.NovoPedidoRepository(db)

	listarPedidosUC := usecase.NovoListarPedidosUseCase(pedidoRepo)

	grpcServer := grpc.NewServer()
	pb.RegisterServicoPedidosServer(
		grpcServer,
		grpcserver.NovoServidorPedidos(listarPedidosUC),
	)

	lis, err := net.Listen("tcp", "50051")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor gRPC: %v", err)
	}
	log.Println("Servidor gRPC iniciado na porta 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Erro ao servir gRPC: %v", err)
	}
}
