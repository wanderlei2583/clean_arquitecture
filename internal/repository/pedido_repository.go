package repository

import "gorm.io/gorm"

type PedidoRepository struct {
	db *gorm.DB
}

func NovoPedidoRepository(db *gorm.DB) *PedidoRepository {
	return &PedidoRepository{
		db: db,
	}
}
