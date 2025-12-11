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


	db,err := repository.NewPostgresDB()
	if err!=nil{
		log.Fatalf("Ошибка подключения к бд:",err)
	}

	storage := repository.NewRoomPostgres(db)
	roomService := service.NewRoomService(storage)
	handler := handler.NewHandler(roomService)

	router := handler.InitRoutes()

	server := http.Server{
		Addr: "0.0.0.0:5050",
		Handler : router,
	}

	fmt.Println("Server started on :5050")
	server.ListenAndServe()

}