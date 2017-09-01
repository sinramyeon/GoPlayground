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

	// 4. Get User Event
	events, _, err := client.Activity.ListEventsPerformedByUser(context.Background(), id, true, nil)
	if err != nil {
		log.Println("[ERROR]", err)
		return false, 1
	}

	for _, v := range events {

		// utctime is user event by utctime.
		utctime := v.GetCreatedAt()
		year, month, day := utctime.In(loc).Date()

		// 5. Find Commit by today, localtime
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
