package main

import (
	"context"
	entity "go_services_lab/pkg/entity"
	handler "go_services_lab/pkg/handler"
	repository "go_services_lab/pkg/repository"
	server "go_services_lab/pkg/server"
	service "go_services_lab/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	ca := cache.New(5*time.Minute, 10*time.Minute)
	ca.Set("tata", 12, cache.DefaultExpiration)
	ca.Set("product1", &entity.Product{1, "Banana", 12.}, cache.DefaultExpiration)
	ca.Set("product2", &entity.Product{2, "Apple", 16.}, cache.DefaultExpiration)
	ca.Set("product3", &entity.Product{3, "Orange", 20.}, cache.DefaultExpiration)
	ca.Set("countProduct", 3, cache.DefaultExpiration)
	ca.Set("order1", &entity.Order{1, 1, entity.Stores{{entity.Product{1, "Banana", 12.}, 10}, {entity.Product{2, "Apple", 16.}, 15}}}, cache.DefaultExpiration)
	ca.Set("order2", &entity.Order{2, 2, entity.Stores{{entity.Product{1, "Banana", 12.}, 2}, {entity.Product{2, "Apple", 16.}, 10}, {entity.Product{3, "Orange", 20.}, 25}}}, cache.DefaultExpiration)
	ca.Set("countOrder", 2, cache.DefaultExpiration)

	c := repository.NewRepository(ca)
	services := service.NewService(c)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(handlers.InitRoutes()); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
