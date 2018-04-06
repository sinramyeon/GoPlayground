package Twitter

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Connect With Twitter.
// I use env.go 's keys
func ConnTwitter() *twitter.Client {

	// 1. Get my auth keys
	var con TwitterConfig
	con = conf(con)

	// 2. you can use oauth1 to set config
	config := oauth1.NewConfig(con.ConfKey, con.ConfSecret)
	token := oauth1.NewToken(con.TokenKey, con.TokenSecret)

	// 3. connect with twitter.
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return client

}

// See my own timeline
func MyTimeline(client *twitter.Client) {

	// you can get your own timeline
	tweets, _, _ := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		// From count, You can set many of parametres.
		// See
		Count: 10,
	})

	for _, v := range tweets {

		//fmt.Println("Coordinates : ", v.Coordinates)
		//fmt.Println("CreatedAt : ", v.CreatedAt)
		//fmt.Println("CurrentUserRetweet : ", v.CurrentUserRetweet)
		//fmt.Println("DisplayTextRange : ", v.DisplayTextRange)
		//fmt.Println("Entities : ", v.Entities)
		//fmt.Println("ExtendedEntities : ", v.ExtendedEntities)
		//fmt.Println("ExtendedTweet : ", v.ExtendedTweet)
		//fmt.Println("FavoriteCount : ", v.FavoriteCount)
		//fmt.Println("Favorited : ", v.Favorited)
		//fmt.Println("FilterLevel : ", v.FilterLevel)
		//fmt.Println("FullText : ", v.FullText)
		//fmt.Println("ID : ", v.ID)
		//fmt.Println("IDStr : ", v.IDStr) // User's Own ID
		//fmt.Println("InReplyToScreenName : ", v.InReplyToScreenName)
		//fmt.Println("InReplyToStatusID : ", v.InReplyToStatusID)
		//fmt.Println("QuotedStatusIDStr : ", v.QuotedStatusIDStr)
		//fmt.Println("RetweetCount : ", v.RetweetCount)
		//fmt.Println("Retweeted : ", v.Retweeted)
		//fmt.Println("RetweetedStatus : ", v.RetweetedStatus)
		//fmt.Println("Scopes : ", v.Scopes)
		//fmt.Println("Source : ", v.Source)
		fmt.Println("Text : ", v.Text) // UserTweet
		//fmt.Println("Truncated : ", v.Truncated)
		//fmt.Println("User : ", v.User)
		fmt.Println("UserName : ", v.User.Name) // UserName
	}

}

// Write Tweet
func SendTweet(client *twitter.Client, str string) {
	client.Statuses.Update(str, nil)
}

// Get Someone's Tweet
func GetUserTweet(client *twitter.Client, name string) {

	log.Println(name + " said.")

	usertweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		Count:           10,
		ScreenName:      name, // @userid
		IncludeRetweets: twitter.Bool(false),
	})

	for _, v := range usertweets {

		fmt.Println(v.Text, nil)

	}
}

// Get Tweet by time
func schedule(client *twitter.Client, delay time.Duration) chan bool {

	// make channel
	stop := make(chan bool)
	go func() {
		for {
			//this goroutine goes foever.
			//and.. show your timeline!
			//of course you can use 'mux' func. it's just for chan study.
			MyTimeline(client)
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

// Almost Same as schedule Method
func scheduleS(client *twitter.Client, delay time.Duration, name string) chan bool {

	stop := make(chan bool)
	go func() {
		for {
			StockingSomebody(client, name)
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

// Just see specific users tweet
func StockingSomebody(client *twitter.Client, name string) {

	// In Your Timeline
	tweets, _, _ := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 10,
	})

	// Detect Specific User
	for _, v := range tweets {

		if strings.Contains(v.User.Name, name) {
			fmt.Println(v.Text)
		}
	}

}

// There is The greatest youtube digger, @XylitoLdrink.
// This func is made for him.

//Get his youtube linked tweet and repost!
func GetXylitolMusic(client *twitter.Client) {

	var URL []string

	// Get his tweet
	usertweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		Count:           20,
		ScreenName:      "XylitoLdrink",
		IncludeRetweets: twitter.Bool(false),
	})

	// Find youtube link
	// All of tweets link looks likt t.sa923skdfl ~~ something.
	// I couldn't find it's youtube or not
	// So sometimes It has error If tweet has ask.fm link, it reposts ask.fm too.
	// I should find the way to detect the link is youtube or not.
	for _, v := range usertweets {

		if strings.Contains(v.Text, "https") {
			URL = append(URL, v.Text)
		}
	}

	// Repost it!
	for _, tweet := range URL {

		client.Statuses.Update(tweet, nil)
	}

}

// Get today's his greatest music link
func TodayXylitol(client *twitter.Client) {

	// for today's tweet
	var TodayTweets []string

	// get what time is it now
	TimeNow := time.Now().Local().Add(time.Hour * time.Duration(-1))
	_, _, Day := TimeNow.Date()

	// get his recent 500 tweets
	// he ussually tweets 100 to 500 in day
	usertweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		Count:           500,
		ScreenName:      "XylitoLdrink",
		IncludeRetweets: twitter.Bool(false),
	})

	for _, v := range usertweets {

		// get when he wrotes this tweet
		TweetTime, _ := v.CreatedAtTime()
		TweetTime.Local()
		_, _, XDay := TweetTime.Date()

		// if the day was today
		if Day == XDay {

			if strings.Contains(v.Text, "https") || strings.Contains(v.Text, "-") {
				// we must get lid of ask.fm links
				// if it was usual tweet, it contains korean. we should get rid of it too.
				r := regexp.MustCompile(`[가-힣]`)
				matches := r.FindAllString(v.Text, -1)

				if len(matches) == 0 {
					TodayTweets = append(TodayTweets, v.Text)
				}
			}
		}
	}

	for _, tweet := range TodayTweets {

		client.Statuses.Update(tweet, nil)
	}

}

// return bool pointer
func Bool(v bool) *bool {
	ptr := new(bool)
	*ptr = v
	return ptr
}

// float to int converter
func FloatToInt(f float64) int {

	var y int = int(f)
	return y
}

// get d-day and tweet!
func D_Day(client *twitter.Client) {

	// 1. get seoul's local time
	timeNow := time.Now().Local()
	seouldate, _ := time.LoadLocation("Asia/Seoul")

	// 2. set due date.
	dueDate := time.Date(int(2017), time.December, int(15), int(0), int(0), int(0), int(0), seouldate)

	// 3. get difference
	diff := timeNow.Sub(dueDate)

	d_day := int(diff.Hours()/24) - 1

	if d_day != 10 || d_day != 9 || d_day != 8 || d_day != 7 || d_day != 6 || d_day != 5 || d_day != 4 || d_day != 3 || d_day != 2 || d_day != 1 || d_day != 0 {

		SendTweet(client, `[알림]
			팔군 @TM_PalGooN 님의 퇴사일이 `+strconv.Itoa(d_day)+"일 남았습니다.")
	} else {

		SendTweet(client, `[경축]
				팔군 @TM_PalGooN 님!
				퇴사일이 `+strconv.Itoa(d_day)+"일 남았습니다. 힘내세요!")
	}
}

// auto search and retweet
func SearchAndRT(client *twitter.Client, s string) {

	// search some tweet with specific words
	search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: s,
		Count: 100,
	})

	// if it contains media or url, and has minimum 5rt, retweet this
	for _, v := range search.Statuses {

		rtcount := v.RetweetCount
		url := len(v.Entities.Urls)
		media := len(v.Entities.Media)

		if rtcount > 5 && (url != 0 || media != 0) {
			client.Statuses.Retweet(v.ID, &twitter.StatusRetweetParams{})
		}

	}

}
