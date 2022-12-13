package main

import (
	"context"
	models "go_services_lab/models"
	handler "go_services_lab/pkg/handler"
	repository "go_services_lab/pkg/repository"
	service "go_services_lab/pkg/service"
	server "go_services_lab/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	cache_order := cache.New(5*time.Minute, 10*time.Minute)
	cache_order.Set("tata", 12, cache.DefaultExpiration)
	cache_order.Set("product1", &models.Product{1, "Banana", 12.}, cache.DefaultExpiration)
	cache_order.Set("product2", &models.Product{2, "Apple", 16.}, cache.DefaultExpiration)
	cache_order.Set("product3", &models.Product{3, "Orange", 20.}, cache.DefaultExpiration)
	cache_order.Set("countProduct", 3, cache.DefaultExpiration)
	cache_order.Set("order1", &models.Order{1, 1, models.Stores{{models.Product{1, "Banana", 12.}, 10}, {models.Product{2, "Apple", 16.}, 15}}}, cache.DefaultExpiration)
	cache_order.Set("order2", &models.Order{2, 2, models.Stores{{models.Product{1, "Banana", 12.}, 2}, {models.Product{2, "Apple", 16.}, 10}, {models.Product{3, "Orange", 20.}, 25}}}, cache.DefaultExpiration)
	cache_order.Set("countOrder", 2, cache.DefaultExpiration)

	repository_order := repository.NewRepositoryOrder(cache_order)
	service_order := service.NewServiceOrder(repository_order)
	handler_order := handler.NewHandlerOrder(service_order)

	server_order := new(server.Server)
	go func() {
		if err := server_order.Run("8000", handler_order.InitRoutesOrder()); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	cache_user := cache.New(5*time.Minute, 10*time.Minute)
	cache_user.Set("user1", &models.User{1, "Alexey", "lewka", "lewka007"}, cache.DefaultExpiration)
	cache_user.Set("user2", &models.User{2, "Ivan", "vane4ka", "trueMan_"}, cache.DefaultExpiration)
	cache_user.Set("user3", &models.User{3, "Masha", "tyan", "mashanyasha"}, cache.DefaultExpiration)
	cache_user.Set("countUser", 3, cache.DefaultExpiration)

	repository_user := repository.NewRepositoryUser(cache_user)
	service_user := service.NewServiceUser(repository_user)
	handler_user := handler.NewHandlerUser(service_user)

	server_user := new(server.Server)
	go func() {
		if err := server_user.Run("8001", handler_user.InitRoutesUser()); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server_order.Shutdown(ctx); err != nil {
		log.Fatal("Server order forced to shutdown:", err)
	}

	if err := server_user.Shutdown(ctx); err != nil {
		log.Fatal("Server user forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
