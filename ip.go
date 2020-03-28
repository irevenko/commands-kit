package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

func getIPConfig() *cobra.Command {
	return &cobra.Command{
		Use:   "all",
		Short: "List of settings",
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
}
