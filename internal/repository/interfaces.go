package repository

import "github.com/Coldwws/todo/internal/models"

type RoomRepository interface {
	GetAllRooms() ([]models.Room,error)
	GetRoomById(id int)(*models.Room,error)
	CreateRoom(room models.Room)(int,string)
	UpdateRoom(id int, updateRoom models.UpdateRoom)(error)
	DeleteRoom(id int)(int,error)
}