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
	URL  string `json:"url"`
}

// PokemonDescription is for detailed info
type PokemonDescription struct {
	CaptureRate int64 `json:"capture_rate"`
	Color       struct {
		Name string `json:"name"`
	} `json:"color"`
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

	resp, err := http.Get(responseObject.Pokemon[randPokemon].Species.URL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var respObj PokemonDescription
	json.Unmarshal(respData, &respObj)

	// total pokemons in the pokedex fmt.Println(len(responseObject.Pokemon))
	fmt.Println(responseObject.Pokemon[randPokemon].Species.Name)
	fmt.Println(respObj.CaptureRate)
	fmt.Println(respObj.Color.Name)
	fmt.Println(responseObject.Pokemon[randPokemon].Species.URL)
}
