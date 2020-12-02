package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "pokemon",
		Short: "Get random pokemon data",
		Long:  "Get random pokemon data from the PokeAPI \nSource code: https://github.com/irevenko/commands-kit",
		Run: func(cmd *cobra.Command, args []string) {
			const url string = "https://pokeapi.co/api/v2/pokedex/national"

			req, err := http.NewRequest("GET", url, nil)
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(body))
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
