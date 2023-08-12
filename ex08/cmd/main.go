package main

import (
	"context"
	"fmt"
	"github.com/rovn208/df-go/ex08/internal/config"
	"github.com/rovn208/df-go/ex08/internal/repo"
	"github.com/rovn208/df-go/ex08/internal/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err = repo.InitializeDB(c.DB_URL); err != nil {
		log.Fatal(err)
	}

	router := routers.SetupRoutes()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", c.Port),
		Handler: router,
	}
	go func() {
		if err = srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Println("Server closed under request")
			} else {
				log.Fatal("Server closed unexpect: ", err)
			}
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}
