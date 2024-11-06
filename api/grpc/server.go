package grpc

import (
	"context"

	"github.com/wanderlei2583/clean_arquitecture/internal/usecase"
)

type ServidorPedidos struct {
	pb.UnimplementedServicePedidosServer
	listarPedidoUseCase *usecase.ListarPedidosUseCase
}

func NovoServidorPedidos(uc *usecase.ListarPedidosUseCase) *ServidorPedidos {
	return &ServidorPedidos{
		listarPedidoUseCase: uc,
	}
}

func (s *ServidorPedidos) ListarPedidos(
	ctx context.Context,
	req *pb.RequisicaoListarPedidos,
) (*pb.RespostaListarPedidos, error) {
	pedidos, err := s.listarPedidoUseCase.Executar(ctx)
	if err != nil {
		return nil, err
	}

	var pbPedidos []*pb.Pedido
	for _, p := range pedidos {
		pbPedidos = append(pbPedidos, &pb.Pedido{
			Id:           p.ID,
			IdCliente:    p.IDCliente,
			Status:       p.Status,
			ValorTotal:   p.ValorTotal,
			CriadoEm:     p.CriadoEm.String(),
			AtualizadoEm: p.AtualizadoEm.String(),
		})
	}

	return &pb.RespostaListarPedidos{
		Pedidos: pbPedidos,
	}, nil
}
