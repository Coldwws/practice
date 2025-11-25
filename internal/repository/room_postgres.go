package repository

import (
	"fmt"

	"github.com/Coldwws/todo/internal/models"
	"github.com/jmoiron/sqlx"
)

type RoomPostgres struct {
	DB *sqlx.DB
}

func NewRoomPostgres(DB *sqlx.DB) *RoomPostgres {
	return &RoomPostgres{
		DB: DB,
	}
}

func(r *RoomPostgres)GetAllRooms()([]models.Room,error){
	var rooms []models.Room
	query := `SELECT * FROM rooms`
	err := r.DB.Select(&rooms,query)
	if err!=nil{
		return nil,err
	}
	return rooms,nil
}


func(r *RoomPostgres)GetRoomById(id int)(*models.Room,error){
	var room models.Room
	query := `SELECT * FROM rooms where id = $1`
	err := r.DB.Get(&room,query,id)
	if err != nil{
		return nil, fmt.Errorf("error : room with id %d not found",id)
	}
	return &room,nil
}

func(r *RoomPostgres)CreateRoom(room models.Room)(int,error){
	var id int
	query := `INSERT INTO rooms(number,type,description) VALUES($1,$2,$3) RETURNING id`
	err := r.DB.QueryRow(query,room.Number,room.Type,room.Description).Scan(&id)
	if err !=nil{
		return 0,fmt.Errorf("Failed to create room : %w",err)
	}
	return id,nil
}

func(r *RoomPostgres)UpdateRoom(id int, updateRoom models.UpdateRoom)(error){
	query := `UPDATE rooms SET number = $1, type = $2, description = $3 WHERE id = $4`
	_,err := r.DB.Exec(query,updateRoom.Number,updateRoom.Type,updateRoom.Description,id)
	if err != nil{
		return fmt.Errorf("Failed to update room: %w",err)
	}
	return nil
}

func(r *RoomPostgres)DeleteRoom(id int)(int,error){
	query := `DELETE FROM rooms WHERE id = $1`
	result,err := r.DB.Exec(query,id)
	if err != nil{
		return 0, fmt.Errorf("Failed to delete room: %w",err)
	}
	rowsAffected,err := result.RowsAffected()
	if err !=nil{
		return 0, fmt.Errorf("Could not get affected rows: %w",err)
	}
	if rowsAffected == 0{
		return 0, fmt.Errorf("Room with id: %d not found",id)
	}
	return id,nil

}