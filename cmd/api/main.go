package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pressly/goose/v3"
	"github.com/SaH1l191/ticket-system/internal/auth"
	"github.com/SaH1l191/ticket-system/internal/config"
	"github.com/SaH1l191/ticket-system/internal/handler"
	"github.com/SaH1l191/ticket-system/internal/repo"
	"github.com/SaH1l191/ticket-system/internal/router"
	"github.com/SaH1l191/ticket-system/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()
	auth.Init(cfg.JWTSecret)

	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	migrationsDir := os.Getenv("MIGRATIONS_DIR")
	if migrationsDir == "" {
		migrationsDir = "./migrations"
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("goose dialect: %v", err)
	}
	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatalf("goose up: %v", err)
	}

	userRepo := repo.NewPostgresUserRepo(db)
	// ticketRepo := repo.NewPostgresTicketRepo(db)

	authSvc := service.NewAuthService(userRepo)
	// ticketSvc := service.NewTicketService(ticketRepo)

	authH := handler.NewAuthHandler(authSvc)
	// ticketH := handler.NewTicketHandler(ticketSvc)

	r := router.New(authH, ticketH)

	srv := &http.Server{
		Addr:              cfg.ServerAddr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		log.Printf("server starting on %s", cfg.ServerAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
	log.Println("server stopped")
}
