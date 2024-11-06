package grpc

import "github.com/wanderlei2583/clean_arquitecture/internal/usecase"

type ServidorPedidos struct {
	pb.UnimplementedServicePedidosServer
	listarPedidoUseCase *usecase.ListarPedidosUseCase
}
