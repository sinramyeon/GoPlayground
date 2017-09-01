package GithubAPI

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	Pushevent        = "PushEvent"
	PullRequestEvent = "PullRequestEvent"
)

func ConnectGithub() (context.Context, *github.Client) {
	// https://developer.github.com/v3/
	// Get your oauth excess token at https://github.com/blog/1509-personal-api-tokens
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		// 3. Set Your Environment Variable
		// Or... You can just put it in here like AccessToken : "blah-blah"

		/*Admonition
		PLEAAAAAAAAAAAAASE DO NOT UPLOAD YOUR OAUTH KEY OR TOKEN TO GITHUB!!!!
		*/
		&oauth2.Token{AccessToken: os.Getenv("GithubAPI")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return ctx, client
}

// Get User Commit
// It returns user commit status with commit count
func GetGitCommit(id string) (b bool, commit int) {

	// 1. Get time
	// I'm Korean, so I got location by Asia/Seoul, But You can change!
	// I made this func for get today's commit status.
	// Fork this and make funny commit bot!
	koryear, kormon, kordate := time.Now().Date()
	loc, _ := time.LoadLocation("Asia/Seoul")

	var commitArray []string

	// 2. Connect With Github
	_, client := ConnectGithub()
	// 3. Get User Event
	events, _, err := client.Activity.ListEventsPerformedByUser(context.Background(), id, true, nil)
	if err != nil {
		log.Println("[ERROR]", err)
		return false, 1
	}

	for _, v := range events {

		// utctime is user event by utctime.
		utctime := v.GetCreatedAt()
		year, month, day := utctime.In(loc).Date()

		// 4. Find Commit by today, localtime
		if ((v.GetType() == Pushevent) || (v.GetType() == PullRequestEvent)) && ((koryear == year) && (kormon == month) && (kordate == day)) {
			commitID := v.GetID()
			commitArray = append(commitArray, commitID)
		}
	}

	// If there is no commit...
	if len(commitArray) == 0 {
		return false, 0
	} else {
		// If there is commit...
		return true, len(commitArray)
	}

}

// GetUserRepo
// It returns your all repository name and description
func GetMyRepo() map[string]string {

	ctx, client := ConnectGithub()

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatal(err)
	}

	repolist := make(map[string]string)

	for _, v := range repos {
		// Name and Description(or sth) is *Name. You should Use Get... Method.
		repolist[v.GetName()] = v.GetDescription()

	}

	return repolist

}

// SearchRepo
// It searches repository with query and returns it's name and description
func SearchRepo(id string) map[string]string {

	ctx, client := ConnectGithub()

	// list all repositories for the authenticated user

	options := github.SearchOptions{
	//put values you need
	}

	repos, _, err := client.Search.Repositories(ctx, id, &options)
	if err != nil {
		log.Fatal(err)
	}

	repolist := make(map[string]string)

	for _, v := range repos.Repositories {
		// Name and Description(or sth) is *Name. You should Use Get... Method.
		repolist[v.GetName()] = v.GetDescription()

	}

	return repolist

}

// GetGist
// It searches repository with user id and returns it's url and description
func GetGist(id string) map[string]string {

	ctx, client := ConnectGithub()

	// list all repositories for the authenticated user

	options := github.GistListOptions{
	//put values you need
	}

	gists, _, err := client.Gists.List(ctx, id, &options)
	if err != nil {
		log.Fatal(err)
	}

	repolist := make(map[string]string)

	for _, v := range gists {
		repolist[v.GetHTMLURL()] = v.GetDescription()

	}
	return repolist

}

// GetUserIdint
// Each user has their own id(int)
func GetUserID(id string) int {

	ctx, client := ConnectGithub()

	var userid int
	options := github.SearchOptions{}
	user, _, err := client.Search.Users(ctx, id, &options)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range user.Users {
		userid = v.GetID()
	}

	for _, v := range user.Users {
		userid = v.GetID()
	}

	return userid

}
