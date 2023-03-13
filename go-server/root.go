package main

import (
	"html/template"
	"log"
	"net/http"
)

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
