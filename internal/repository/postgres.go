package repository

import (
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB() (*sqlx.DB,error){
	
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_NAME"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
	)

	db,err := sqlx.Connect("postgres",dsn)
	if err !=nil{
		return nil,err
	}

	return db,err
}
