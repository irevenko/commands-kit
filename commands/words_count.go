package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "words_count",
		Short: "Counts words number",
		Long:  "An util for counting words number \nSource code: https://github.com/irevenko/commands-kit",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := ioutil.ReadFile(args[0])
			if err != nil {
				fmt.Println("File reading error", err)
				return
			}

			wordsCount := strings.Fields(string(data))

			fmt.Println(len(wordsCount))
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
