package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "rm",
		Short: "Delete files and dirs",
		Long:  "and dirs \nSource code: https://github.com/irevenko/commands-kit",
		Run: func(cmd *cobra.Command, args []string) {
			err := os.Remove(args[0])
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
