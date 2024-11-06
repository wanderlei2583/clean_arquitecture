package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/wanderlei2583/clean_arquitecture/internal/domain"
)

type PedidoRepository struct {
	db *gorm.DB
}

func NovoPedidoRepository(db *gorm.DB) *PedidoRepository {
	return &PedidoRepository{
		db: db,
	}
}

func (r *PedidoRepository) BuscarTodos(
	ctx context.Context,
) ([]domain.Pedido, error) {
	var pedidos []domain.Pedido
	result := r.db.WithContext(ctx).Find(&pedidos)
	if result.Error != nil {
		return nil, result.Error
	}
	return pedidos, nil
}
