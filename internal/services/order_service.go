package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/Louis-Ai/insurance-order-batcher/internal/models"
	"github.com/Louis-Ai/insurance-order-batcher/internal/utility"
)

type OrderService struct {
	outputDirectory string
	batchSize       int
	currentBatch    []models.Order
	mu              sync.Mutex
}

func NewOrderService(outputDirectory string, batchSize int) *OrderService {
	return &OrderService{
		outputDirectory: outputDirectory,
		batchSize:       batchSize,
		currentBatch:    make([]models.Order, 0),
	}
}

func (s *OrderService) AddOrderToBatch(order *models.Order) error {
	err := utility.ValidatePostcode(order.Address.Postcode)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.currentBatch = append(s.currentBatch, *order)

	if len(s.currentBatch) >= s.batchSize {
		err := s.writeBatchToFile(&s.currentBatch)
		if err != nil {
			return fmt.Errorf("error writing batch to file: %v ", err)
		}
		s.currentBatch = make([]models.Order, 0)
	}

	return nil

}

func (s *OrderService) writeBatchToFile(orders *[]models.Order) error {
	filename := fmt.Sprintf("%s/orders_%d.csv", s.outputDirectory, time.Now().Unix())
	fmt.Printf("Writing batch to: %s\n", filename)

	return utility.WriteToCSV(filename, orders)
}
