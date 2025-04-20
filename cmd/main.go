package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Louis-Ai/insurance-order-batcher/internal/config"
	"github.com/Louis-Ai/insurance-order-batcher/internal/handlers"
	"github.com/Louis-Ai/insurance-order-batcher/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}

	orderService := services.NewOrderService(cfg.OutputDirectory, cfg.BatchSize)

	orderHandler := handlers.NewOrderHandler(orderService)

	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		orderHandler.SubmitOrder(w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server listening on port %s \n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
