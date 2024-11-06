package domain

import "time"

type Pedido struct {
	ID           string    `json:"id"            gorm:"primary_key"`
	ClienteID    string    `json:"id_cliente"    gorm:"column:id_cliente"`
	Status       string    `json:"status"`
	ValorTotal   float64   `json:"valor_total"   gorm:"column:valor_total"`
	CriadoEm     time.Time `json:"criado_em"     gorm:"column:criado_em"`
	AtualizadoEm time.Time `json:"atualizado_em" gorm:"column:atualizado_em"`
}

func (Pedido) TableName() string {
	return "pedidos"
}
