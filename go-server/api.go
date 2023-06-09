package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func api(writer http.ResponseWriter, request *http.Request) {
	responseData := make(map[string]string)

	responseData["timestamp"] = timestamp()
	for name, values := range request.Header {
		if inSlice(strings.ToLower(name), []string{"x-real-ip", "x-forwarded-for"}) {
			continue
		}
		for _, value := range values {
			responseData[name] = value
			log.Println("responseData - ", name, ": ", value)
		}
	}

	// get DA data
	daData := getData(responseData)
	log.Println(daData)

	// add DA data to the response map
	for prop, val := range daData["properties"] {
		responseData[prop] = val
	}

	// DB access seems a bit slow, it can be done
	// concurrently with sending the response
	var wg sync.WaitGroup

	// log data to the DB
	go func() {
		wg.Add(1)
		defer wg.Done()

		// prepare a special map for DB
		// ~~the data in it has lowercased headers, better keep that~~
		DBEntries := make(map[string]string)
		log.Println("preparing log entries")
		for k, v := range responseData {
			//DBEntries[strings.ToLower(k)] = v
			DBEntries[k] = v
		}
		for failCounter := 0; failCounter < 5; failCounter++ {
			log.Printf("db attempt %d", failCounter)
			key, err := db.Put(DBEntries)
			if err != nil {
				log.Printf("can't store data in DB: %s, attempt: %d", err, failCounter)
			} else {
				log.Printf("stored %s in the db, attempt %d\n", key, failCounter)
				break
			}
			time.Sleep(1 / 100)
		}
	}()

	// send response to the client
	writer.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		log.Println("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = writer.Write(responseJSON)
	if err != nil {
		log.Println("can't send a response", err)
	} else {
		log.Println("response sent")
	}

	// make sure the threads finish before this function does
	wg.Wait()
}
