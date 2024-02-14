package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/viciousvs/currencyAPI/config"
	"github.com/viciousvs/currencyAPI/internal/repository"
	httpServer "github.com/viciousvs/currencyAPI/internal/server/http"
	"github.com/viciousvs/currencyAPI/internal/server/http/routes"
	"github.com/viciousvs/currencyAPI/internal/storage"
	"github.com/viciousvs/currencyAPI/internal/usecase"
)

func init() {
	if err := godotenv.Load("dev.env"); err != nil {
		log.Printf("cannot load dev.env file=> %v", err)
		log.Println("used default values for config")
	}
}

func main() {
	cfg := config.NewConfig()
	s := new(httpServer.Server)
	db := storage.NewPostgresDB(cfg.PostgresConfig)
	CurrencyRepo := repository.NewCurrencyRepo(db)
	CurrencyUseCase := usecase.NewCurrencyUseCase(CurrencyRepo)

	h := routes.NewHandler(CurrencyUseCase)
	mux := h.InitRoutes()

	go func() {
		// service connections
		if err := s.Run(cfg.ServerConfig, mux); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
