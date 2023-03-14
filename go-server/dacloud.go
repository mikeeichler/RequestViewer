package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getData(headers map[string]string) {
	url := "http://region2.deviceatlascloud.com:80/v1/detect/properties?licencekey="
	licenseKey := os.Getenv("DA_LICENSE")
	url += licenseKey
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("can't send request to da servers", err)
	}
	for h, v := range headers {
		ch := clientHints()
		if inSlice(h, ch) || h == "User-Agent" {
			req.Header.Add("X-DA-"+h, v)
		} else if strings.EqualFold(h, "cookie") {
			c := dapropsFromCookie(v)
			req.Header.Add("X-DA-Client-Properties", c)
		} else {
			req.Header.Add(h, v)
		}
	}
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("didn't get e response from DA cloud", err)
	}
	log.Println(url)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("unreadable response", err)
	}
	fmt.Println(body)
}
