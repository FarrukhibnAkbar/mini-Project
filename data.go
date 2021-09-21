package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	_"github.com/lib/pq"
)

var(
	Host="localhost"
	Username="admin"
	Phassword=12345
	Database="login"
	Port=5432
)

type Password struct{
	Id int
	UserLogin string
	UserPassword string
}

type PasswordBody struct {
	UserLogin string
	UserPassword string
}

var SQL_GET_LOGIN =`
select
	login_id,
	login,
	password
from Login;
`

var SQL_POST_LOGIN =`
insert into Login(login, password)values($1,$2)
returning
	login_id,
	login,
	password
`


func GetLogin() []Password {

	c := fmt.Sprintf(
		"host=%s user=%s password=%d dbname=%s port=%d",
		Host, Username, Phassword, Database, Port,
	)
	
	db, err := sql.Open("postgres", c)

	if err != nil{
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query(SQL_GET_LOGIN)

	defer rows.Close()

	if err != nil{
		panic(err)
	}

	var passwords []Password

	for rows.Next(){

		var password Password

		rows.Scan(&password.Id, &password.UserLogin, &password.UserPassword)

		passwords = append(passwords, password)
	}

	return passwords
}



func PostLogin(body PasswordBody) Password {

	c := fmt.Sprintf(
		"host=%s user=%s password=%d dbname=%s port=%d",
		Host, Username, Phassword, Database, Port,
	)

	db, err := sql.Open("postgres",c)

	defer db.Close()

	if err != nil{
		panic(err)
	}

	var newPassword Password

	err = db.QueryRow(SQL_POST_LOGIN, body.UserLogin, body.UserPassword).
	Scan(&newPassword.Id, &newPassword.UserLogin, &newPassword.UserPassword)

	if err != nil{
		panic(err)
	}

	return newPassword
}


func GetAdmin(w http.ResponseWriter, r *http.Request){

	e := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	e.Encode(GetLogin())
}

func PostAdmin(w http.ResponseWriter, r *http.Request){

	e := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")
	
	var newPassword PasswordBody

	body,_ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &newPassword)

	e.Encode(PostLogin(newPassword))
}













