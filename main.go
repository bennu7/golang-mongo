package main

import (
	"fmt"
	"net/http"

	"github.com/bennu7/golang-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user", uc.GetAllUser)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017/")
	if err != nil {
		fmt.Println("status panic err connect localhost => ", err)
		panic(err)
	}
	fmt.Println("status connect localhost", s)
	return s
}
