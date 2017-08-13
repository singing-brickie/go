package main

import (
	"net/http"
)


func checkErr(err error) {
			if err != nil {
					panic(err)
			}
}




func main() {
  fs := http.FileServer(http.Dir("../client-side/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/pint_fine", pint_fine)
	http.HandleFunc("/login", login)
	http.HandleFunc("/", login)


	http.ListenAndServe(":8080", nil)
}
