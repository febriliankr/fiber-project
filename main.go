package main

import (
	"encoding/json"
	"fmt"
	"go-mysql-react/helpers"
	"go-mysql-react/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users models.Users
	var arrUser []models.Users
	var response models.Response

	db := helpers.ConnectDB()
	defer db.Close()

	rows, err := db.Query("Select id, name ,email, password from users")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.Name, &users.Email, &users.Password); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/getUsers", returnAllUsers).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	fmt.Println("Connected to http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
