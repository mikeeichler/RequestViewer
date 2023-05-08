package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

func uaViewer(writer http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Subpath: %s\n", path)
	requestData := make(map[string]string)
	requestData["message"] = strings.TrimPrefix(request.URL.Path, "/ua_viewer/")
	log.Println("message " + requestData["message"])
	var requestString string
	tmpl, err := template.ParseFiles("./templates/ua_viewer.html")
	if err != nil {
		log.Fatal("can't parse the template", err)
	}

	// prepare the request to be sent to the API for logging (in case of previewers and similar, where there is not JS)
	client := &http.Client{}
	apiURL := "https://rv.mikee.site/api"
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Clacks-Overhead", "GNU Terry Pratchett")
	req.Header.Set("X-Message", requestData["message"])
	for name, values := range request.Header {
		// Loop over all values for the name.
		clientHintSlice := clientHints()
		ch := fmt.Sprintf(strings.Join(clientHintSlice[:], ","))
		writer.Header().Set("Accept-CH", ch)
		writer.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
		req.Header.Set("Accept-CH", ch)
		for _, value := range values {
			// fmt.Println(name, value)
			// writer.Write([]byte(fmt.Sprintf("%s: %s\n", name, value)))
			req.Header.Set(name, value)
			requestData[name] = value
			requestString = fmt.Sprintf("%s, '%s': '%s'", requestString, name, value)
			if name == "User-Agent" {
				requestData["ua"] += value
			}
		}
	}
	// send it to the API for needs logging
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		var resp *http.Response
		resp, err = client.Do(req)
		counter := 1
		for err != nil {
			log.Printf("can't send data to the API %v, attempt: %d", err, counter)
			resp, err = client.Do(req)
			counter += 1
			if counter > 5 {
				log.Println("final attempt to send data to the API failed")
			}
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println("resp body close failed")
			}
		}(resp.Body)
	}()
	requestData["requestString"] = requestString
	err = tmpl.Execute(writer, requestData)
	if err != nil {
		return
	}
}
