package main

import (
	"GoVCS/information"
	"GoVCS/mySecrets"
	"context"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func main() {
	token := mySecrets.TOKEN
	owner := mySecrets.USERNAME
	repoName := "GoVCS"

	// Create GitHub client
	ctx := context.Background() // Learn more about context package
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// func GetRepoInformation(client *github.Client, ctx context.Context, owner, repoName string) {
	information.GetRepoInformation(client, ctx, owner, repoName)
}
