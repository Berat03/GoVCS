package information

import (
	"context"
	"fmt"
	"github.com/google/go-github/v39/github"
	"os"
	"reflect"
)

func derefPointer(value reflect.Value) interface{} {
	if value.Kind() == reflect.Ptr && !value.IsNil() {
		return value.Elem().Interface()
	}
	return value.Interface()
}

func GetRepoInformation(client *github.Client, ctx context.Context, owner, repoName string) {
	repository, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		fmt.Printf("Error fetching repo information: %v\n", err)
		os.Exit(1)
	}

	repositoryValue := reflect.ValueOf(repository).Elem()
	repositoryType := repositoryValue.Type()

	for i := 0; i < repositoryValue.NumField(); i++ {
		field := repositoryValue.Field(i)
		fieldName := repositoryType.Field(i).Name
		fieldValue := derefPointer(field)
		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}
}
