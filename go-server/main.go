package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	port := os.Getenv("PORT")
	httpServer := &http.Server{
		// Addr:    "127.0.0.1:5550",
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),
		Handler: mux,
	}
	mux.HandleFunc("/", root)
	if err := httpServer.ListenAndServe(); err == http.ErrServerClosed {
		log.Println("Web server closed")
	} else {
		log.Fatalln(err)
	}

}

func root(writer http.ResponseWriter, request *http.Request) {
	for name, values := range request.Header {
		// Loop over all values for the name.
		writer.Header().Set("Accept-CH", "Sec-CH-UA, Sec-CH-UA-Arch, Sec-CH-UA-Bitness, Sec-CH-UA-Full-Version-List, Sec-CH-UA-Full-Version, Sec-CH-UA-Mobile, Sec-CH-UA-Model, Sec-CH-UA-Platform, Sec-CH-UA-Platform-Version")
		writer.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
		for _, value := range values {
			// fmt.Println(name, value)
			writer.Write([]byte(fmt.Sprintf("%s: %s\n", name, value)))
		}
		// writer.WriteHeader(200)
	}

}
