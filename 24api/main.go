package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Anime struct {
	ID       int     `json:"animeid"`
	Name     string  `json:"name"`
	Episodes int     `json:"episodes"`
	Author   *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Twitter  string `json:"twitter"`
}

// fake database
var animes []Anime

// middleware
func (a *Anime) isEmpty() bool {
	// return a.ID == 0 && a.Name == ""
	return a.Name == ""
}

// seeding of fake db
func main() {
	fmt.Println("Hello from the Anime API !")
	r := mux.NewRouter()

	animes = append(animes, Anime{ID: 101, Name: "One Piece", Episodes: 1016,
		Author: &Author{Fullname: "Eiichiro Oda", Twitter: "@EiichiroOda"}})

	animes = append(animes, Anime{ID: 102, Name: "Black Clover", Episodes: 300,
		Author: &Author{Fullname: "Yuki Tabata", Twitter: "@TabataNoYuki"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/animes", getAllAnimes).Methods("GET")
	r.HandleFunc("/anime/{id}", getAnime).Methods("GET")
	r.HandleFunc("/anime", createAnime).Methods("POST")
	r.HandleFunc("/anime/{id}", updateAnime).Methods("PUT")
	r.HandleFunc("/anime/{id}", deleteAnime).Methods("DELETE")

	//listening
	log.Fatal(http.ListenAndServe(":8080", r))

}

// controllers - file/folders
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Anime API</h1>"))
}

func getAllAnimes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All Animes")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animes)
}

func getAnime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Anime")
	w.Header().Set("Content-Type", "application/json")

	// take id from req
	params := mux.Vars(r)

	// convert params[id] to an integer
	id, _ := strconv.Atoi(params["id"])

	// loop through animes, find matching id and return the anime
	for _, anime := range animes {
		if anime.ID == id {
			json.NewEncoder(w).Encode(anime)
			return
		}
	}
	json.NewEncoder(w).Encode("No anime found with that ID")
	return
}

func createAnime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add an Anime")
	w.Header().Set("Content-Type", "application/json")

	//base condition
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// decode the request body
	var anime Anime
	_ = json.NewDecoder(r.Body).Decode(&anime)
	if anime.isEmpty() {
		json.NewEncoder(w).Encode("No data found")
		return
	}

	//generate unique id
	rand.Seed(time.Now().UnixNano())
	anime.ID = rand.Intn(100)

	// add the new anime to the fake db
	animes = append(animes, anime)
	json.NewEncoder(w).Encode(anime) // auto exit the method
}

func updateAnime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updated an Anime")
	w.Header().Set("Content-Type", "application/json")

	//base condition
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// take id from req
	params := mux.Vars(r)

	// convert params[id] to an integer
	id, _ := strconv.Atoi(params["id"])

	//loop
	for index, anime := range animes {
		if anime.ID == id {
			animes = append(animes[:index], animes[index+1:]...)
			var updateAnime Anime
			_ = json.NewDecoder(r.Body).Decode(&updateAnime)
			updateAnime.ID = id
			animes = append(animes, updateAnime)
			json.NewEncoder(w).Encode(updateAnime)
		}
	}
}

func deleteAnime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an Anime")
	w.Header().Set("Content-Type", "application/json")

	// take id from req
	params := mux.Vars(r)

	// convert params[id] to an integer
	id, _ := strconv.Atoi(params["id"])

	//loop
	for index, anime := range animes {
		if anime.ID == id {
			animes = append(animes[:index], animes[index+1:]...)
			json.NewEncoder(w).Encode("Anime deleted")
		}
	}
}
