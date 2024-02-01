package main

import (
	"GoVCS/mySecrets"
	"context"
	"fmt"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
	"os"
	"reflect"
)

func derefPointer(value reflect.Value) interface{} {
	if value.Kind() == reflect.Ptr && !value.IsNil() {
		return value.Elem().Interface()
	}
	return value.Interface()
}

func main() {
	// Set your GitHub token
	token := mySecrets.TOKEN

	// Create a GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Specify the repository owner and name
	owner := mySecrets.USERNAME
	repoName := "GoVCS"

	// Retrieve repository information
	repository, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		fmt.Printf("Error fetching repository information: %v\n", err)
		os.Exit(1)
	}

	// Iterate over the fields of the repository object and print them
	repositoryValue := reflect.ValueOf(repository).Elem()
	repositoryType := repositoryValue.Type()
	fmt.Println("Repository Information:")
	for i := 0; i < repositoryValue.NumField(); i++ {
		field := repositoryValue.Field(i)
		fieldName := repositoryType.Field(i).Name
		fieldValue := derefPointer(field)
		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}

}
