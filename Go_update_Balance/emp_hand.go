package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.First(&employee)
	json.NewEncoder(w).Encode(employee)

}

func updateb(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.First(&employee, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(employee)

}
