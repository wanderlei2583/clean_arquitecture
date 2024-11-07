package main

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/wanderlei2583/clean_arquitecture/api/grpc"
	"github.com/wanderlei2583/clean_arquitecture/api/grpc/pb"
	"github.com/wanderlei2583/clean_arquitecture/api/rest"
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

	router := gin.Default()
	handlerPedidos := rest.NovoHandlerPedidos(listarPedidosUC)
	rest.SetupRoutes(router, handlerPedidos)

	grpcServer := grpc.NewServer()
	pb.RegisterServicoPedidosServer(
		grpcServer,
		grpc.NovoServidorPedidos(listarPedidosUC),
	)

	go func() {
		lis, err := net.Listen("tcp", "50051")
		if err != nil {
			log.Fatalf("Erro ao iniciar o servidor gRPC: %v", err)
		}
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Erro ao servir gRPC: %v", err)
		}
	}()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor REST: %v", err)
	}
}
