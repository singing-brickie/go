package main

import (
  "fmt"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
  "encoding/json"
  "io/ioutil"
)

type dbDetails struct {
	Password string `json:"password"`
	Host string `json:"host"`
  User string `json:"user"`
}

func getDetails() dbDetails {
  plan, _ := ioutil.ReadFile("dbConnectionDetails.json")
  var data dbDetails
  err := json.Unmarshal(plan, &data)
  checkErr(err)
  return data
}


var data dbDetails = getDetails()

func addLogin(userName string, password string) {
	hash, _ := HashPassword(password)
	db, err := sql.Open("mysql", data.User + ":" + data.Password + data.Host + "pint_fine")
	checkErr(err)
	stmt, err := db.Prepare("INSERT users SET userName='"+userName+"' ,password='"+hash+"'")
  checkErr(err)
  res,err := stmt.Exec()
  checkErr(err)
  affect, err := res.RowsAffected()
  checkErr(err)
  fmt.Println(affect)
	//panic(res)
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
	db, err := sql.Open("mysql", data.User + ":" + data.Password + data.Host + "pint_fine")
	checkErr(err)
	rows, err := db.Query("SELECT password FROM users WHERE userName="+"'" + userName +"'")
  checkErr(err)
  db.Close();
  for rows.Next() {
	  var db_password string
    err = rows.Scan(&db_password)
    checkErr(err)
    return CheckPasswordHash(password,db_password)
  }
  return false
}
