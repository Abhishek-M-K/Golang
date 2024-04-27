package main

import (
	"encoding/json"
	"fmt"
)

type manga struct {
	Name   string `json:"mangaName"`
	Genre  string
	Author string
	Price  int
	Status string
	Tags   []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON in GoLang")
	encodeJson()
	decodeJson()
}

func encodeJson() {
	mangaLibrary := []manga{
		{"One Piece", "Adventure", "Eiichiro Oda", 100, "Ongoing", []string{"Pirates", "Action", "Comedy", "Adventure"}},
		{"Berserk", "Action", "Kentaro Miura", 150, "Ongoing", []string{"Dark Fantasy", "Horror", "Action"}},
		{"Naruto", "Adventure", "Masashi Kishimoto", 120, "Completed", []string{"Ninja", "Action", "Adventure"}},
		{"Attack on Titan", "Action", "Hajime Isayama", 130, "Completed", nil},
	}

	// package this data into a JSON format

	// finalJson, err := json.Marshal(mangaLibrary) // Marshal takes an interface and returns a byte slice
	finalJson, err := json.MarshalIndent(mangaLibrary, "", "  ") // MarshalIndent takes an interface and returns a byte slice with indentation

	CheckNilError(err)

	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonData := []byte(`
			{
				"mangaName": "One Piece",
				"Genre": "Adventure",
				"Author": "Eiichiro Oda",
				"Price": 100,
				"Status": "Ongoing",
				"tags": ["Pirates", "Action", "Comedy", "Adventure"]
			}
	`)

	var mangaData manga

	isValid := json.Valid(jsonData)

	if isValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &mangaData) // Unmarshal takes a byte slice and a pointer to an interface
		// dont send the struct directly, send the pointer to the struct instead because Unmarshal needs to modify the struct

		fmt.Printf("%#v\n", mangaData)
	}else{
		fmt.Println("JSON is invalid")
	}

	// some cases where we just want to add data to key value pairs
	fmt.Println("-------------------------Decoding JSON to map---------------------------")

	var data map[string]interface{}
	json.Unmarshal(jsonData, &data)
	fmt.Printf("%#v\n", data)

	for k, v := range data {
		fmt.Printf("Key: %v, Value: %v\n, Type:%T\n", k, v, v)
	}

}

func CheckNilError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
