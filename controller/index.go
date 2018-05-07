package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/articles/{category}", categoryHandler)

	r.HandleFunc("/user", userAdditionHandler).Methods("POST")
	r.HandleFunc("/user", userListHandler).Methods("GET")
	r.HandleFunc("/relationship", relationshipListAddHandler).Methods("GET")
	r.HandleFunc("/relationship", relationshipUpdateAddHandler).Methods("PUT")

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

