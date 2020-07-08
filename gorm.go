package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

type User struct {
	ID       int64 `gorm:"primary_key"`
	Username string
}

func (s *User) TableName() string {
	return "auth_user"
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "user"
	PASS := "password"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "dbname"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/users", GetAllUsers),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
	db := gormConnect()
	defer db.Close()

	var allUsers []User
	db.Find(&allUsers)
	fmt.Println(allUsers)

	w.WriteHeader(http.StatusOK)
	w.WriteJson(&allUsers)
}
