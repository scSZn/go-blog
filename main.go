package main

import (
	"github.com/scSZn/blog/conf"
	"github.com/scSZn/blog/internal/routers"
	"log"
	"net/http"
)

func main() {
	Init()
	engine := routers.NewRouter()
	server := http.Server{
		Handler: engine,
		Addr:    conf.GetListenAddr(),
	}
	log.Fatal(server.ListenAndServe())
}
