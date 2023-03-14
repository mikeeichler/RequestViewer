package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func root(writer http.ResponseWriter, request *http.Request) {
	requestData := make(map[string]string)
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Fatal("can't parse the template", err)
	}
	for name, values := range request.Header {
		// Loop over all values for the name.
		clientHintSlice := clientHints()
		ch := fmt.Sprintf(strings.Join(clientHintSlice[:], ","))
		writer.Header().Set("Accept-CH", ch)
		writer.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
		for _, value := range values {
			// fmt.Println(name, value)
			// writer.Write([]byte(fmt.Sprintf("%s: %s\n", name, value)))
			requestData[name] = value
		}
	}
	tmpl.Execute(writer, requestData)
}
