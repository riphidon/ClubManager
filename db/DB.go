package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/riphidon/clubmanager/config"
)

var DB *sql.DB

func InitDB() {
	connectDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		config.Data.Host, config.Data.DBPort, config.Data.DBUser, config.Data.DBPass, config.Data.DBName)
	var err error
	DB, err = sql.Open("postgres", connectDB)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}
