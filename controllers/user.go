package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bennu7/golang-mongo/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Println("id => ", id)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	fmt.Println("oid => ", oid)

	u := models.User{}

	err := uc.session.DB("go_api").C("users").FindId(oid).One(&u)
	if err != nil {
		fmt.Println("err connect database name go_api => ", err)
		w.WriteHeader(404)
		return
	}
	// {
	// 	w.WriteHeader("db go_api not found in mongodb",http.StatusNotFound)
	// 	return
	//  }

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error in marshalling json", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println(u)

	u.Id = bson.NewObjectId()

	uc.session.DB("go_api").C("users").Insert(u)

	json.Marshal(u)
}
func DeleteUser() {

}
