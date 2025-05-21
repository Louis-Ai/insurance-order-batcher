package utility

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/Louis-Ai/insurance-order-batcher/internal/models"
)

//todo: add functionality for kafka to add new entries for each order
//todo: db connection to also store orders rather than in memory store

func WriteToCSV(filename string, orders *[]models.Order) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"CustomerID", "Address", "Timestamp"}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	for _, order := range *orders {
		address := fmt.Sprintf("%v %v %v", order.Address.AddressLineOne, order.Address.TownCity, order.Address.Postcode)
		timestamp := order.OrderTime.String()

		row := []string{order.CustomerID, address, timestamp}

		err = writer.Write(row)
		if err != nil {
			return err
		}
	}
	return nil

}
