package handler

import (
	"net/http"
	"github.com/Coldwws/todo/internal/auth"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) login(c *gin.Context){
		var input LoginInput
		if err := c.ShouldBindJSON(&input);err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID := 1
		token,err := auth.GenerateToken(userID)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
}