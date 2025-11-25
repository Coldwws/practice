package service

import "github.com/Coldwws/todo/internal/models"

type RoomService interface {
	GetAllRooms() ([]models.Room,error)
	GetRoomById(id int)(*models.Room,error)
	CreateRoom(room models.Room)(int,string)
	UpdateRoom(id int,update models.UpdateRoom)(error)
	DeleteRoom(id int)(int,error)
}