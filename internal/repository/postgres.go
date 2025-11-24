package repository

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB() (*sqlx.DB,error){
	
	const(
		host = "localhost"
		port =  5432
		user = "postgres"
		password = "Edil3875319"
		dbname = "practice-todo"
	)

	dsn := fmt.Sprintf("host=%s port =%d user =%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)

	db,err := sqlx.Connect("postgres",dsn)
	if err !=nil{
		return nil,err
	}

	return db,err
}
