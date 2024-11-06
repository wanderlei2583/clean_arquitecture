package main

import (
	"log"
	"net"

	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=db_pedidos port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	pedidoRepo := repository.NovoPedidoRepository(db)

	listarPedidosUC := usecase.NovoListarPedidosUseCase(pedidoRepo)

	router := gin.Default()
	handlerPedidos := rest.NovoHandlerPedidos(listarPedidosUC)
	rest.SetupRoutes(router, handlerPedidos)

	grpcServer := grpc.NewServer()
	pb.RegisterServicePedidosServer(
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
