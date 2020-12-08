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
		Use:   "chars_count",
		Short: "Counts chars number",
		Long:  "An util for counting chars number \nSource code: https://github.com/irevenko/commands-kit",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := ioutil.ReadFile(args[0])
			if err != nil {
				fmt.Println("File reading error", err)
				return
			}

			charsCount := strings.Split(string(data), "")

			fmt.Println(len(charsCount))
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
