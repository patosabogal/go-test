package service

import (
	"errors"
	"github.com/patosabogal/go-test/src/domain"
)

var users map[string]*domain.User
var tweets map[int]*domain.Tweet
var id int
var logged string

//TODO: GET TWIT BY USER Y DELETE USER
func InitializeService() {
	tweets = make(map[int]*domain.Tweet)
	users = make(map[string]*domain.User)
	id = 0
	logged = ""
}

func AddUser(username string, password string) error {
	if _, ok := users[username]; ok {
		err := errors.New("Username already exists")
		return err
	}
	newUser := domain.NewUser(username, password)
	users[username] = newUser
	return nil
}

func 

func LogIn(username string, password string) error {
	if _, ok := users[username]; !ok {
		if password != users[username].Password {
			err := errors.New("Username or password not valid")
			return err
		}
	}
	return nil
}

func PublishTweet(t *domain.Tweet) (int, error) {
	if logged == "" {
		err := errors.New("You must log in to be able to publish a tweet")
		return -1, err
	}
	if t.Text == "" {
		err := errors.New("text is required")
		return -1, err
	}
	username := t.Username
	users[username].Tweets = append(users[username].Tweets, *t)
	tweets[id] = t
	oldId := id
	id = id + 1
	return oldId, nil
}

func GetTweets() map[int]*domain.Tweet {

	return tweets
}

func GetTweetById(i int) (*domain.Tweet, error) {

	if 0 <= i && i < id {
		return tweets[i], nil
	}
	err := errors.New("id not valid")
	return nil, err

}
