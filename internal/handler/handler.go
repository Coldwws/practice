package handler

import (
	"github.com/Coldwws/todo/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.RoomService
}

func NewHandler(service service.RoomService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler)InitRoutes() *gin.Engine{
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	room := router.Group("/room")
	{
		room.GET("/",h.GetAllRooms)
		room.GET("/:id",h.GetRoomById)
		room.POST("/create",h.CreateRoom)
		room.PATCH("/update/:id",h.UpdateRoom)
		room.DELETE("/delete/:id",h.DeleteRoom)
	}

	return router
}
