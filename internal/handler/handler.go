package handler

import (
	"github.com/Coldwws/todo/internal/auth"
	"github.com/Coldwws/todo/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.RoomService
	auth 	service.AuthService
}

func NewHandler(service service.RoomService, auth service.AuthService) *Handler {
	return &Handler{
		service: service,
		auth:    auth,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", h.login)
		authRoutes.POST("/register", h.register)
	}
	

	room := router.Group("/room")
	room.Use(auth.AuthRequired())
	{
		room.GET("/", h.GetAllRooms)
		room.GET("/:id", h.GetRoomById)
		room.POST("/create", h.CreateRoom)
		room.PATCH("/update/:id", h.UpdateRoom)
		room.DELETE("/delete/:id", h.DeleteRoom)
	}

	return router
}
