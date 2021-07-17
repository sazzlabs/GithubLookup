package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/go-github/v37/github"
)


func main() {
	githubUsername, _ := requestGithubUsername()
	userInfo, _ := userLookup(*githubUsername)

	twitterURL := "https://twitter.com/" + *userInfo.TwitterUsername

	fmt.Println("Name:", *userInfo.Name)
	fmt.Println("Bio:", *userInfo.Bio)
	fmt.Println("Location:", *userInfo.Location)
	fmt.Println("Public Repos:", *userInfo.PublicRepos)
	fmt.Println("Public Gists:", *userInfo.PublicGists)
	fmt.Println("Followers:", *userInfo.Followers)
	fmt.Println("Following:", *userInfo.Following)
	fmt.Println("Twitter:", twitterURL)
	fmt.Println("Profile:", *userInfo.HTMLURL)
	fmt.Println("Created at:", *userInfo.CreatedAt)
}

func requestGithubUsername() (*string, error) {
	var githubUsername string
	fmt.Print("Github Username: ")

	_, err := fmt.Scanln(&githubUsername)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	return &githubUsername, nil
}

func userLookup(githubUsername string) (*github.User, error) {
	client := github.NewClient(nil)
	user, _, err := client.Users.Get(context.Background(), githubUsername)

	
	if err != nil {
		fmt.Println(errors.New("error in fetching the user"))
		fmt.Println(err)
		os.Exit(1)
	}

	return user, err
}