package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "pwd",
		Short: "Get current working dir",
		Long:  "An util for getting current working dir \nSource code: https://github.com/irevenko/commands-kit",
		Run: func(cmd *cobra.Command, args []string) {
			path, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(path)
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
