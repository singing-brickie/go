package main

import (
	"html/template"
	"net/http"
	"fmt"


)

type ContactDetails struct {
	userName   string
	password string
}


func main() {


  fs := http.FileServer(http.Dir("../client-side/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	tmpl := template.Must(template.ParseFiles("../client-side/static/html/index.html"))


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Println("in if")
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			userName:   r.FormValue("user"),
			password: r.FormValue("pass"),
		}

		fmt.Println(details.userName)

	})

	http.ListenAndServe(":8080", nil)
}
