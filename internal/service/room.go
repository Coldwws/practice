package service

import (
	"github.com/Coldwws/todo/internal/models"
	"github.com/Coldwws/todo/internal/repository"
)

type roomService struct {
	storage repository.RoomRepository
}

func NewRoomService(storage repository.RoomRepository) RoomService {
	return &roomService{storage: storage}
}


func(s *roomService)GetAllRooms()([]models.Room,error){
	rooms ,_ := s.storage.GetAllRooms()

	return rooms,nil
}

func (s *roomService)GetRoomById(id int)(*models.Room,error){
	return s.storage.GetRoomById(id)
}

func (s *roomService)CreateRoom(room models.Room)(int,string){
	return s.storage.CreateRoom(room)
}

func(s *roomService)UpdateRoom(id int,update models.UpdateRoom)error{
	return  s.storage.UpdateRoom(id,update)
}

func(s *roomService)DeleteRoom(id int)(int,error){
	return s.storage.DeleteRoom(id)
}