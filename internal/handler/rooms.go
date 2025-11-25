package handler

import (
	"net/http"
	"strconv"

	"github.com/Coldwws/todo/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllRooms(c *gin.Context) {
	rooms, err := h.service.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

func (h *Handler) GetRoomById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	room, err := h.service.GetRoomById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"room": room,
	})

}

func (h *Handler) CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBind(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	id, errStr := h.service.CreateRoom(room)
	if errStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})

}

func (h *Handler) UpdateRoom(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),})
			return
	}

	var update models.UpdateRoom
	if err := c.ShouldBind(&update);err != nil{
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	err = h.service.UpdateRoom(id,update)
	if err != nil{
		c.JSON(http.StatusNotFound,err.Error())
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Room updated succesfully"})

}

func (h *Handler) DeleteRoom(c *gin.Context){
	idStr := c.Param("id")
	id,err := strconv.Atoi(idStr)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	_, err = h.service.DeleteRoom(id)
	c.JSON(http.StatusOK,gin.H{
		"message":"Room deleted successfully",
	})
}