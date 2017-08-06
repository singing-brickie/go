package main

import (
	"html/template"
	"net/http"
)



func main() {
    fs := http.FileServer(http.Dir("../client-side/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	tmpl := template.Must(template.ParseFiles("../client-side/static/html/todos.html"))


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}
