package repository

import (
	"GoVCS/mySecrets"
	"context"
	"fmt"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
	"log"
)

// CreateGitHubRepository creates a new GitHub repository with the given client and repository details.
func CreateRepository() error {
	token := mySecrets.TOKEN
	userName := mySecrets.USERNAME

	ctx := context.Background() // Learn more about context package
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	description := "My GitHub repository" // Repository description
	private := false                      // Whether the repository is private
	autoInit := true                      // Whether to create an initial commit with an empty README

	r := &github.Repository{Name: &userName, Private: &private, Description: &description, AutoInit: &autoInit}
	repo, _, err := client.Repositories.Create(ctx, userName, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
	return nil
}
