package main

import (
	"html/template"
	"net/http"
	"fmt"
)


func pint_fine(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", 301)
	} else {
		if r.Method != http.MethodPost {
		  tmpl := template.Must(template.ParseFiles("../client-side/static/html/secret.html"))
		  tmpl.Execute(w, nil)
			 return
		}
		details := ContactDetails{
			userName:   r.FormValue("user"),
			password: r.FormValue("pass"),
		}
		addLogin(details.userName,details.password)
		fmt.Println("success")
	}
}
