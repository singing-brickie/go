package main

import (
	"html/template"
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"


)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

type ContactDetails struct {
	userName   string
	password string
}

func checkErr(err error) {
			if err != nil {
					panic(err)
			}
}

func addLogin(userName string, password string) {
	hash, _ := HashPassword(password)
	db, err := sql.Open("mysql", "root:shreebo1@tcp(localhost:3306)/pint_fine")
	checkErr(err)
	stmt, err := db.Prepare("INSERT users SET userName='"+userName+"' ,password='"+hash+"'")
  checkErr(err)
  res, err := stmt.Exec()
  checkErr(err)
  id, err := res.LastInsertId()
  checkErr(err)
  fmt.Println(id)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func queryLogin(userName string, password string) bool {
	db, err := sql.Open("mysql", "££££££££££££££££") // route to database
	checkErr(err)
	rows, err := db.Query("SELECT password FROM users WHERE userName="+"'" + userName +"'")
  checkErr(err)
  for rows.Next() {
		fmt.Println("in")
	  var db_password string
    err = rows.Scan(&db_password)
    checkErr(err)
    return CheckPasswordHash(password,db_password)
  }
  return false
}

func secret(w http.ResponseWriter, r *http.Request) {
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
	  http.Redirect(w, r, "/secret", 301)
	} else {
		http.Redirect(w, r, "/login", 301)
	}
}

func main() {
  fs := http.FileServer(http.Dir("../client-side/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)


	http.ListenAndServe(":8080", nil)
}
