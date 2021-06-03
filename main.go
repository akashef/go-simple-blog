package main

import (
	"log"
	"net/http"

	_ "github.com/gorilla/mux"

	"github.com/rs/cors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func DBDropMigrate() {
	db.DropTableIfExists(&User{}, &Post{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
}
func DBForeignKey() {
	db.Model(&Post{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	db.Model(&Comment{}).AddForeignKey("post_id", "post(id)", "CASCADE", "CASCADE")
	db.Model(&Comment{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
}
func DBConnection() {
	dbinfo := "host=localhost user=postgres password=postgres dbname=360imaging port=5432 sslmode=disable"
	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		panic("failed to connect database")
	}
	DBDropMigrate()
	DBForeignKey()
}

func main() {
	DBConnection()
	router := routerHandler()
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
