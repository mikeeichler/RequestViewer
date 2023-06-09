package main

import (
	"fmt"
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"log"
	"net/http"
	"os"
)

// Globals

var db *base.Base

func main() {
	detaKey := os.Getenv("RV_DB_KEY")
	d, err := deta.New(deta.WithProjectKey(detaKey))
	if err != nil {
		log.Println("can't connect to the DB")
	}
	db, err = base.New(d, "logs")
	if err != nil {
		log.Println("failed to init new Base instance:", err)
	}
	mux := http.NewServeMux()
	port := os.Getenv("PORT")
	log.Println("$PORT is: ", port)
	httpServer := &http.Server{
		// Addr:    "127.0.0.1:5550",
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),
		Handler: mux,
	}
	mux.HandleFunc("/api", api)
	mux.HandleFunc("/logs", logs)
	mux.HandleFunc("/ua_viewer/", uaViewer)
	//mux.HandleFunc("/ua_viewer/{message:([a-zA-Z0-9=\\-\\/])*}", uaViewer)
	mux.HandleFunc("/", root)
	// this enables serving JavaScript and CSS
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/static/"))))
	if err := httpServer.ListenAndServe(); err == http.ErrServerClosed {
		log.Println("Web server closed")
	} else {
		log.Println(err)
	}

}
