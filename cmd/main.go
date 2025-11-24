package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Coldwws/todo/internal/handler"
	"github.com/Coldwws/todo/internal/repository"
	"github.com/joho/godotenv"
)




func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env: %v", err)
	}
	mux := http.NewServeMux()
	db,err := repository.NewPostgresDB()
	if err!=nil{
		log.Fatalf("Ошибка подключения к бд:",err)
	}

	roomHandler := handler.NewRoomHandler(repository.NewRoomPostgres(db))
	mux.HandleFunc("/",roomHandler.GetRoom)
	// mux.HandleFunc("/room/create",roomHandler.CreateRoom)
	// mux.HandleFunc("/room/delete",roomHandler.DeleteRoom)
	// mux.HandleFunc("/room/update",roomHandler.UpdateRoom)

	server := http.Server{
		Addr: ":5050",
		Handler : mux,
	}

	fmt.Println("Server started on :5050")
	server.ListenAndServe()

}