package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// A Pokedex struct to map the Entire Response
type Pokedex struct {
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int `json:"entry_number"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}

// PokemonDescription is for detailed info
type PokemonDescription struct {
	CaptureRate       int `json:"capture_rate"`
	BaseHappiness     int `json:"base_happiness"`
	FlavorTextEntries []struct {
		Text     string `json:"flavor_text"`
		Language struct {
			Name string `json:"name"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
	GrowthRate struct {
		Name string `json:"name"`
	} `json:"growth_rate"`
	Color struct {
		Name string `json:"name"`
	} `json:"color"`
	Shape struct {
		Name string `json:"name"`
	} `json:"shape"`
}

func main() {
	var cmd = &cobra.Command{
		Use:   "pokemon",
		Short: "Get Random Pokemon Data",
		Long:  "A simple CLI tool for getting random pokemon data \nSource code: https://github.com/irevenko/commands-kit",
		Run: func(cmd *cobra.Command, args []string) {
			rand.Seed(time.Now().UnixNano())

			pokedexResp, err := http.Get("http://pokeapi.co/api/v2/pokedex/national/")
			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			pokedexData, err := ioutil.ReadAll(pokedexResp.Body)
			if err != nil {
				log.Fatal(err)
			}

			var pokedexObject Pokedex
			json.Unmarshal(pokedexData, &pokedexObject)

			randPokemon := rand.Intn(len(pokedexObject.Pokemon) + 1)

			pokemonResp, err := http.Get(pokedexObject.Pokemon[randPokemon].Species.URL)
			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			pokemonData, err := ioutil.ReadAll(pokemonResp.Body)
			if err != nil {
				log.Fatal(err)
			}

			var pokemonObject PokemonDescription
			json.Unmarshal(pokemonData, &pokemonObject)

			fmt.Println("Name: " + pokedexObject.Pokemon[randPokemon].Species.Name)
			fmt.Println("Capture rate: " + strconv.Itoa(pokemonObject.CaptureRate))
			fmt.Println("Base Happiness: " + strconv.Itoa(pokemonObject.BaseHappiness))
			fmt.Println("Pokedex Number: " + strconv.Itoa(pokedexObject.Pokemon[randPokemon].EntryNo))
			fmt.Println("Color: " + pokemonObject.Color.Name)
			fmt.Println("Shape: " + pokemonObject.Shape.Name)
			fmt.Println("Growth Rate: " + pokemonObject.GrowthRate.Name)
			for k, v := range pokemonObject.FlavorTextEntries {
				if v.Language.Name == "en" {
					fmt.Println("Flavor Text: " + pokemonObject.FlavorTextEntries[k].Text)
					break
				}
			}
			fmt.Println(pokedexObject.Pokemon[randPokemon].Species.URL)
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
