package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"encoding/json"
	"github.com/WTIFS/tantan-demo/model"
	"github.com/WTIFS/tantan-demo/service"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/users").Subrouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/articles/{category}", categoryHandler)

	s.HandleFunc("/", userAdditionHandler).Methods("POST")
	s.HandleFunc("/", userListHandler).Methods("GET")
	s.HandleFunc("/{user_id}/relationships", relationshipListAddHandler).Methods("GET")
	s.HandleFunc("/{user_id}/relationships/{other_user_id}", relationshipUpdateAddHandler).Methods("PUT")

	serverMsg := http.ListenAndServe(":8000", r)
	log.Fatal(serverMsg)


}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "param: %v", vars["category"])
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello!\n"))
	//fmt.Fprintf(w, "Hello!")
}

//add user
/**
POST /users
Create a user
allowed fields:
 name = string
Example:
$curl -XPOST -d '{"name":"Alice"}' "http://localhost:80/users"
{
 "id": "11231244213",
 "name": "Alice" ,
 "type": "user"
}
 */
func userAdditionHandler(w http.ResponseWriter, r *http.Request) {
	//name := r.FormValue("name")
	//mobile := r.FormValue("mobile")
	//u := &model.User{
	//	Name: name,
	//	Mobile: mobile,
	//}
	var u *model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if _, err := service.AddUser(u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, u)
}


//list users
/**
GET /users
List all users
Example:
$curl -XGET "http://localhost:80/users"
[
 {
 "id": "21341231231",
 "name": "Bob" ,
 "type": "user"
 },
 {
 "id": "31231242322",
 "name": "Samantha" ,
 "type": "user"
 }
]
 */
func userListHandler(w http.ResponseWriter, r *http.Request) {
	userList, err := service.ListUsers()
	if (err != nil) {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, userList)
	}
}


//list relationship
/**
Relationships:
GET /users/:user_id/relationships
List a users all relationships
Example:
$curl -XGET "http://localhost:80/users/11231244213/relationships"
[
 {
 "user_id": "222333444",
 "state": "liked" ,
 "type": "relationship"
 },
 {
 "user_id": "333222444",
 "state": "matched" ,
 "type": "relationship"
 },
 {
 "user_id": "444333222",
 "state": "disliked" ,
 "type": "relationship"
 }
]
 */
func relationshipListAddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["user_id"]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if (err != nil) {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		relationList, err:= service.ListRelationsByUserId(userId)
		if (err != nil) {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			respondWithJSON(w, http.StatusOK, relationList)
		}
	}
}


//put relationship
/**
PUT /users/:user_id/relationships/:other_user_id
Create/update relationship state to another user.
allowed fields:
 state = "liked"|"disliked"
If two users have "liked" each other, then the state of the relationship is "matched"
Example:
$curl -XPUT -d '{"state":"liked"}'
"http://localhost:80/users/11231244213/relationships/21341231231"
{
 "user_id": "21341231231",
 "state": "liked" ,
 "type": "relationship"
}
$curl -XPUT -d '{"state":"liked"}'
"http://localhost:80/users/21341231231/relationships/11231244213"
{
 "user_id": "11231244213",
 "state": "matched" ,
 "type": "relationship"
}
$curl -XPUT -d '{"state":"disliked"}'
"http://localhost:80/users/21341231231/relationships/11231244213"
{
 "user_id": "11231244213",
 "state": "disliked" ,
 "type": "relationship"
}
 */
func relationshipUpdateAddHandler(w http.ResponseWriter, r *http.Request) {

}


func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
