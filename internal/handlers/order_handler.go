package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Louis-Ai/insurance-order-batcher/internal/models"
	"github.com/Louis-Ai/insurance-order-batcher/internal/services"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(os *services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: os,
	}
}

func (h *OrderHandler) SubmitOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	err = h.orderService.AddOrderToBatch(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
