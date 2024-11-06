package usecase

import (
	"context"

	"github.com/wanderlei2583/clean_arquitecture/internal/domain"
)

type PedidoRepository interface {
	BuscarTodos(ctx context.Context) ([]domain.Pedido, error)
}

type ListarPedidosUseCase struct {
	repo PedidoRepository
}

func NovoListarPedidosUseCase(repo PedidoRepository) *ListarPedidosUseCase {
	return &ListarPedidosUseCase{
		repo: repo,
	}
}

func (uc *ListarPedidosUseCase) Executar(
	ctx context.Context,
) ([]domain.Pedido, error) {
	return uc.repo.BuscarTodos(ctx)
}
