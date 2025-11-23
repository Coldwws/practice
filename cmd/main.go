package main

import (
	"fmt"
	"net/http"

	"github.com/Coldwws/todo/internal/handler"
)




func main() {
	mux := http.NewServeMux()
	
	roomHandler := handler.NewRoomHandler()
	mux.HandleFunc("/",roomHandler.GetRoom)
	mux.HandleFunc("/room/create",roomHandler.CreateRoom)
	mux.HandleFunc("/room/delete",roomHandler.DeleteRoom)
	mux.HandleFunc("/room/update",roomHandler.UpdateRoom)

	server := http.Server{
		Addr: ":5050",
		Handler : mux,
	}

	fmt.Println("Server started on :5050")
	server.ListenAndServe()

}