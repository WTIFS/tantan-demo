package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/WTIFS/tantan-demo/service"
	"github.com/WTIFS/tantan-demo/model"
	"encoding/json"
	"strconv"
)

func HandlerUser (r *mux.Route) {
	s := r.Subrouter()
	s.HandleFunc("", userAdditionHandler).Methods("POST")
	s.HandleFunc("", userListHandler).Methods("GET")
	s.HandleFunc("/{user_id}/relationships", relationshipListHandler).Methods("GET")
	s.HandleFunc("/{user_id}/relationships/{other_user_id}", relationshipAddHandler).Methods("PUT")
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
	var u *model.User

	//check request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if (u.Name == "") {
		respondWithError(w, http.StatusBadRequest, "Params is empty")
		return
	}

	//add user
	if _, err := service.AddUser(u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	u.SetType()
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
		for i := range userList {
			userList[i].SetType()
		}
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
func relationshipListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["user_id"]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if (err != nil) {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		relationshipList, err:= service.ListRelationsByUserId(userId)
		if (err != nil) {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			for i := range relationshipList {
				relationshipList[i].SetType()
			}
			respondWithJSON(w, http.StatusOK, relationshipList)
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
func relationshipAddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["user_id"]
	otherUserIdStr := vars["other_user_id"]
	userId, err1 := strconv.ParseInt(userIdStr, 10, 64)
	otherUserId, err2 := strconv.ParseInt(otherUserIdStr, 10, 64)
	if (err1 != nil) {
		respondWithError(w, http.StatusInternalServerError, err1.Error())
		return
	} else if (err2 != nil) {
		respondWithError(w, http.StatusInternalServerError, err2.Error())
		return
	}
	var relationship *model.Relationship
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&relationship); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if (relationship.State == "") {
		respondWithError(w, http.StatusBadRequest, "Params is empty")
		return
	}
	defer r.Body.Close()
	relationship.FromUserId = userId
	relationship.ToUserId = otherUserId

	res, err3 := service.UpsertRelationship(relationship)
	if (err3 != nil) {
		panic(err3)
		respondWithError(w, http.StatusBadRequest, err3.Error())
		return
	}
	res.SetType()
	respondWithJSON(w, http.StatusOK, res)

}