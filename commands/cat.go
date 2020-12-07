package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "cat",
		Short: "Implimentation of linux's cat",
		Long:  "Usage: cat <YOUR FILE.TXT>\nSource code: https://github.com/irevenko/commands-kit",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := ioutil.ReadFile(args[0])
			if err != nil {
				fmt.Println("File reading error", err)
				return
			}

			fmt.Println(string(data))
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
