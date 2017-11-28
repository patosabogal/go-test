package service_test

import (
	"github.com/patosabogal/go-test/src/domain"
	"github.com/patosabogal/go-test/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	// t.Fatal("not implemented")
	service.InitializeService()
	user := "grupoesfera"
	text := "This is my first tweet"

	tweet := domain.NewTweet(user, text)
	service.PublishTweet(tweet)

	publishedTweet := service.GetTweets()[0]

	if publishedTweet.Username != user && publishedTweet.Text != text {

		t.Error("Expected tweet is %s: %s \nbut is  %s: %s", user, text, publishedTweet.Username, publishedTweet.Text)

	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	//Initialization
	service.InitializeService()
	var tweet *domain.Tweet

	user := "grupoesfera"

	tweet = domain.NewTweet(user, "")

	_, err := service.PublishTweet(tweet)

	//Validation

	if err == nil {

		t.Error("Expected error")
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")

	}
}

func TestTweetsSlice(t *testing.T) {
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := *publishedTweets[0]
	secondPublishedTweet := *publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, user, text) {
		return
	}
	if !isValidTweet(t, secondPublishedTweet, user, secondText) {
		return
	}

}

func isValidTweet(t *testing.T, twit domain.Tweet, us string, tex string) bool {
	if twit.Username != us && twit.Text != tex {
		t.Errorf("Expected tweet with user %s, and text %s", us, tex)
		return false
	}
	return true
}

func TestCanRetrieveTweetById(t *testing.T) {

	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int
	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	id, _ = service.PublishTweet(tweet)

	publishedTweet, _ := service.GetTweetById(id)

	isValidTweet(t, *publishedTweet, user, text)

}
