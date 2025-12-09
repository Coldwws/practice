package repository

import (
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB() (*sqlx.DB,error){
	
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	db,err := sqlx.Connect("postgres",dsn)
	if err !=nil{
		return nil,err
	}

	return db,err
}
