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

func (uc UserController) GetAllUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("go to get all user")
	var users []models.User

	err := uc.session.DB("go_api").C("users").Find(nil).All(&users)
	if err != nil {
		fmt.Println("error in get all user", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		return
	}

	uj, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error in marshalling json", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Println("get by id => ", id)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	fmt.Println("oid => ", oid)

	u := models.User{}

	err := uc.session.DB("go_api").C("users").FindId(oid).One(&u)
	if err != nil {
		fmt.Println("err find id => ", err)
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error in marshalling json", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println(u)

	u.Id = bson.NewObjectId()

	uc.session.DB("go_api").C("users").Insert(u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error in marshalling json", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Println("go to delete id => ", id)

	if bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	err := uc.session.DB("go_api").C("users").RemoveId(oid)
	if err != nil {
		fmt.Println("error in delete user", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted user %s", oid)
}
