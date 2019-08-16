package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/riphidon/clubmanager/config"
	"github.com/riphidon/clubmanager/db"
	"github.com/riphidon/clubmanager/server"
)

func main() {
	//logger := log.New(os.Stdout, "club", log.LstdFlags|log.Lshortfile)
	//h := server.NewHandler(logger)
	var err error
	config.Data.ParseConfigFile()
	db.InitDB()

	mux := http.NewServeMux()
	addr := config.Data.ServerPort
	srv := server.New(mux, addr)

	server.SetupRoutes(mux)
	fmt.Println(db.AllUsers())
	err = srv.ListenAndServe()
	log.Fatal(err)
}
