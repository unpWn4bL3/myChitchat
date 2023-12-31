package main

import (
	"log"
	"net/http"

	"github.com/unpWn4bl3/myChitchat/config"
	"github.com/unpWn4bl3/myChitchat/routes"
)

func main() {
	startWebServer("8080")
}

func startWebServer(port string) {
	config := config.LoadConfig()
	r := routes.NewRouter()

	// 将静态资源/static存储在/public
	assests := http.FileServer(http.Dir(config.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assests))

	// assets := http.FileServer(http.Dir("./test"))
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)
	log.Println("Starting HTTP service at " + config.App.Address)
	err := http.ListenAndServe(config.App.Address, nil) // Goroutine will block here
	if err != nil {
		log.Println("An error occured starting HTTP listener at " + config.App.Address)
		log.Println("Error: " + err.Error())
	}
}
