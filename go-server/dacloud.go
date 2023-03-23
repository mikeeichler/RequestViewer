package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getData(headers map[string]string) (m map[string]map[string]string) {
	url := "http://region2.deviceatlascloud.com:80/v1/detect/properties?licencekey="
	licenseKey := os.Getenv("DA_LICENSE")
	url += licenseKey
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("can't send request to da servers", err)
	}
	// req.Header.Set("X-DA-Client-Properties", headers["Daprops"])
	req.Header["X-DA-Client-Properties"] = []string{headers["Daprops"]}
	log.Println("X-DA-Client-Properties: '", req.Header["X-DA-Client-Properties"], "'")
	req.Header.Add("Accept", "*/*")
	for h, v := range headers {
		// ch := clientHints()
		if strings.ToLower(h) != "cookie" {
			req.Header["X-DA-"+h] = []string{v}
		}
	}
	log.Println("headers sent to cloud: ", req.Header)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("didn't get e response from DA cloud", err)
	}
	defer resp.Body.Close()
	m = make(map[string]map[string]string)
	// https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
	json.NewDecoder(resp.Body).Decode(&m)
	log.Println("respones map: ", m)
	return m
}
