package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	model "./model" // example of import package
	"github.com/gorilla/mux"
)

var people []model.Person // imported from model

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	people = append(people, model.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &model.Address{City: "City X", State: "State X"}})
	people = append(people, model.Person{ID: "2", Firstname: "Kohn", Lastname: "Foe", Address: &model.Address{City: "City X", State: "State X"}})
	people = append(people, model.Person{ID: "3", Firstname: "Lohn", Lastname: "Goe"})

	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetPeople to get lists of people.
func GetPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(people)
}

// GetPerson to get person respective to the id.
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(&model.Person{})
	//r = mux.Vars(r)
}

// CreatePerson to add person.
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	/*
		To print the response to string
	*/
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	/*
		Parse JSON object without struct
	*/
	m := map[string]interface{}{}
	err := json.Unmarshal(bodyBytes, &m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
	fmt.Println(m["firstname"])

	json.Unmarshal(bodyBytes, &person) // parse JSON to person struct object
	fmt.Println(person.Firstname)
	people = append(people, person)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(people)
}
