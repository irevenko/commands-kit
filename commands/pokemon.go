package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// PokemonSpecies is the Struct to map Pokemon Struct
type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randPokemon := rand.Intn(899)

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/national/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(len(responseObject.Pokemon))
	fmt.Println(responseObject.Pokemon[randPokemon].Species.Name)

}
