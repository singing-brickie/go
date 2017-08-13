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

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	tmpl := template.Must(template.ParseFiles("../client-side/static/html/index.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := ContactDetails{
		userName:   r.FormValue("user"),
		password: r.FormValue("pass"),
	}

  fmt.Println(queryLogin(details.userName,details.password))
	if (queryLogin(details.userName,details.password)) {
		session.Values["authenticated"] = true
	  session.Save(r, w)
	  http.Redirect(w, r, "/pint_fine", 301)
	} else {
		http.Redirect(w, r, "/login", 301)
	}
}
