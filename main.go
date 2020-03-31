package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "myip",
		Short:        "Check your IP Adress Data",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			const url string = "https://ipinfo.io"

			req, _ := http.NewRequest("GET", url, nil)
			res, _ := http.DefaultClient.Do(req)
			body, _ := ioutil.ReadAll(res.Body)
			defer res.Body.Close()

			fmt.Println(string(body))
			return nil
		},
	}

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
