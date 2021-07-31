package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Name string
}

var (
	homeTemplate *template.Template
	contactTemplate *template.Template
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	err := homeTemplate.Execute(w, &Data{
		Name: "Kirill",
	})

	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to 404 page</h1>")
}

func main() {

	homeTemplate = parseTemplate("home.gohtml")
	contactTemplate = parseTemplate("contact.gohtml")

	router := httprouter.New()

	router.GET("/", home)
	router.GET("/contact", contact)

	router.NotFound = http.HandlerFunc(notFound)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func parseTemplate(path string) *template.Template {
	t, err := template.ParseFiles("views/" + path)

	if err != nil {
		panic(err)
	}

	return t
}



