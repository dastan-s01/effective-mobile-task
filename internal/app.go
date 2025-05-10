package internal

import (
	httpSwagger "github.com/swaggo/http-swagger"

	"context"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"taskEffectiveMobile/db"
	"taskEffectiveMobile/internal/app"
	"taskEffectiveMobile/internal/app/handlers"
	"time"
)

// @title Effective Mobile API
// @version 1.0
// @description This is a test server for Effective Mobile
// @host localhost:8080
// @BasePath /

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbPool, err := db.PgConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	di := app.NewDI(dbPool)
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	handlers.InitRoutes(r, di)

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}
	go func() {
		log.Println("Сервер запущен на порту", os.Getenv("PORT"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска сервера: %s\n", err)
		}
	}()

	// Graceful Shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Завершение работы сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении: %v", err)
	}

	log.Println("Сервер остановлен корректно")
}
