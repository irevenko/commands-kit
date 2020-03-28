package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "myip",
		Short:        "Check your IP Adress Data",
		SilenceUsage: true,
	}

	cmd.AddCommand(getIPConfig())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
