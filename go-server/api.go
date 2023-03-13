package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

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
	_, err = writer.Write(responseJSON)
	if err != nil {
		log.Fatal("can't send a response", err)
	}
}
