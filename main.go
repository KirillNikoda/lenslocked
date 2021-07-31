package main

import (
	"fmt"
	"net/http"
)



func main() {
	http.HandleFunc("/", handlerFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site23!</h1>")
}