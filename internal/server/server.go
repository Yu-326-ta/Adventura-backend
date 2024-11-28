package server

import (
	"Adventura/config"
	"Adventura/internal/dao"
	"Adventura/internal/infrastracture"
	"Adventura/internal/routes"
	"Adventura/internal/usecase"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func StartServer() error {
	db, err := infrastracture.NewDB(config.MySQLConfig())
	if err != nil {
		return err
	}
	defer db.Close()
	// Ping the database to ensure connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}
	log.Println("Successfully connected to the database")

	addr := ":" + strconv.Itoa(config.Port())
	if os.Getenv("ENV") == "production" {
		addr = "0.0.0.0" + addr
	}
	log.Printf("Serve on http://%s", addr)

	HealthUsecase := usecase.NewHealth(dao.NewHealth())

	r := routes.NewRouter(
		HealthUsecase,
	)
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := srv.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	return nil
}
