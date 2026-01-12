package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Coldwws/todo/internal/handler"
	"github.com/Coldwws/todo/internal/repository"
	"github.com/Coldwws/todo/internal/service"
)

func main() {
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к бд: %v", err)
	}

	// repositories
	roomRepo := repository.NewRoomPostgres(db)
	userRepo := repository.NewUserPostgres(db)

	// services
	roomService := service.NewRoomService(roomRepo)
	authService := service.NewAuthService(userRepo)

	// handlers
	h := handler.NewHandler(roomService, authService)
	router := h.InitRoutes()

	server := http.Server{
		Addr:    "0.0.0.0:5050",
		Handler: router,
	}

	fmt.Println("Server started on :5050")
	log.Fatal(server.ListenAndServe())
}
