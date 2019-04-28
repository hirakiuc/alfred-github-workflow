package main

import (
	"fmt"
	"os"

	"github.com/hirakiuc/alfred-github-workflow/internal/api"
)

func main() {
	client := api.GetClient()

	/*
		client.FetchReposByUserWithHandler("hirakiuc", func(repos []*github.Repository, err error, hasNext bool) bool {
			if err != nil {
				fmt.Printf("Error occurred: %v", err)
				return false
			}

			for _, repo := range repos {
				fmt.Printf("%s\n", repo.GetFullName())
			}

			return hasNext
		})
	*/

	pulls, err := client.FetchPullRequests("hirakiuc", "tinybucket")
	if err != nil {
		fmt.Printf("Error occurred: %v", err)
		os.Exit(1)
	}

	for _, pull := range pulls {
		fmt.Printf("%s : %s\n", pull.GetURL(), pull.GetTitle())
	}
}
