package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

// Globals

var db *base.Base

func main() {
	deta_key := os.Getenv("RV_KEY")
	d, err := deta.New(deta.WithProjectKey(deta_key))
	if err != nil {
		log.Fatal("can't connect to the DB")
	}
	db, err = base.New(d, "logs")
	if err != nil {
		log.Fatal("failed to init new Base instance:", err)
	}
	mux := http.NewServeMux()
	port := os.Getenv("PORT")
	httpServer := &http.Server{
		// Addr:    "127.0.0.1:5550",
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),
		Handler: mux,
	}
	mux.HandleFunc("/", root)
	mux.HandleFunc("/api", api)
	// this enables serving JavaScript and CSS
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/static/"))))
	if err := httpServer.ListenAndServe(); err == http.ErrServerClosed {
		log.Println("Web server closed")
	} else {
		log.Fatalln(err)
	}

}

func root(writer http.ResponseWriter, request *http.Request) {
	requestData := make(map[string]string)
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Fatal("can't parse the template", err)
	}
	for name, values := range request.Header {
		// Loop over all values for the name.
		writer.Header().Set("Accept-CH", "Sec-CH-UA, Sec-CH-UA-Arch, Sec-CH-UA-Bitness, Sec-CH-UA-Full-Version-List, Sec-CH-UA-Full-Version, Sec-CH-UA-Mobile, Sec-CH-UA-Model, Sec-CH-UA-Platform, Sec-CH-UA-Platform-Version")
		writer.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
		for _, value := range values {
			// fmt.Println(name, value)
			// writer.Write([]byte(fmt.Sprintf("%s: %s\n", name, value)))
			requestData[name] = value
		}
	}
	tmpl.Execute(writer, requestData)
}

func timestamp() (timestamp string) {
	currentTime := time.Now()
	yr := fmt.Sprintf("%04d", currentTime.Year())
	mo := fmt.Sprintf("%02d", int(currentTime.Month()))
	dy := fmt.Sprintf("%02d", currentTime.Day())
	hr := fmt.Sprintf("%02d", currentTime.Hour())
	mi := fmt.Sprintf("%02d", currentTime.Minute())
	sc := fmt.Sprintf("%02d", currentTime.Second())
	timestamp = fmt.Sprintf("%s-%s-%s_T_%s:%s:%s", yr, mo, dy, hr, mi, sc)
	return

}
func api(writer http.ResponseWriter, request *http.Request) {
	responseData := make(map[string]string)
	writer.Header().Set("Content-Type", "application/json")

	responseData["timestamp"] = timestamp()
	for name, values := range request.Header {
		for _, value := range values {
			responseData[name] = value
		}
	}
	responseJSON, err := json.Marshal(responseData)
	// prepare a special map for DB
	// the data in it has lowercased headers, better keep that
	DBEntries := make(map[string]string)
	for k, v := range responseData {
		DBEntries[strings.ToLower(k)] = v
	}
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	key, err := db.Put(DBEntries)
	if err != nil {
		log.Fatal("can't store data in dB", err)
	} else {
		log.Printf("stored %s in the db\n", key)
	}
	writer.Write(responseJSON)
}
