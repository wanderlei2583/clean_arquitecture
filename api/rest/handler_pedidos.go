package rest

import "github.com/wanderlei2583/clean_arquitecture/internal/usecase"

type HandlerPedidos struct {
	ListarPedidosUseCase *usecase.ListarPedidosUseCase
}

func NovoHandlerPedidos(uc *usecase.ListarPedidosUseCase) *HandlerPedidos {
	return &HandlerPedidos{
		ListarPedidosUseCase: uc,
	}
}
