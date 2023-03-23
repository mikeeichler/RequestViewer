package main

import (
	"encoding/json"
	"github.com/deta/deta-go/service/base"
	"log"
	"net/http"
)

func logs(writer http.ResponseWriter, request *http.Request) {
	var data []interface{}
	query := base.Query{
		{},
	}
	_, err := db.Fetch(
		&base.FetchInput{
			Q:    query,
			Dest: &data,
		},
	)
	if err != nil {
		log.Fatal("couldn't fetch data from db, ", err)
	}

	writer.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal("couldn't create response JSON from data, ", err)
	}
	_, err = writer.Write(responseJSON)
	if err != nil {
		log.Fatal("couldn't send response, ", err)
	}
}
