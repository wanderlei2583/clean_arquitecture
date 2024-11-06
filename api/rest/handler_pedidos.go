package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wanderlei2583/clean_arquitecture/internal/usecase"
)

type HandlerPedidos struct {
	ListarPedidosUseCase *usecase.ListarPedidosUseCase
}

func NovoHandlerPedidos(uc *usecase.ListarPedidosUseCase) *HandlerPedidos {
	return &HandlerPedidos{
		ListarPedidosUseCase: uc,
	}
}

func (h *HandlerPedidos) ListarPedidos(c *gin.Context) {
	pedidos, err := h.ListarPedidosUseCase.Executar(c.Request.Context())
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Erro ao buscar pedidos: " + err.Error()},
		)
		return
	}
	c.JSON(http.StatusOK, pedidos)
}

func SetupRoutes(r *gin.Engine, handler *HandlerPedidos) {
	r.GET("/pedidos", handler.ListarPedidos)
}
