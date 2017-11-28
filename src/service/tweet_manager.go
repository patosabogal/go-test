package service

import (
	"errors"
	"github.com/patosabogal/go-test/src/domain"
)

var users map[string]*domain.User
var tweets map[int]*domain.Tweet
var id int
var logged string

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

func DeleteUser(username string, password string) error {
	if _, ok := users[username]; !ok {
		err := errors.New("Username or password not valid")
		return err
	}
	if password != users[username].Password {
		err := errors.New("Username or password not valid")
		return err
	}
	delete(users, username)
	return nil
}

func GetTwitsByUser(username string) ([]domain.Tweet, error) {
	if _, ok := users[username]; !ok {
		err := errors.New("Username not valid")
		return nil, err
	}
	twee := users[username].Tweets
	return twee, nil

}

func LogIn(username string, password string) error {
	if _, ok := users[username]; !ok {
		err := errors.New("Username or password not valid")
		return err
	}
	if password != users[username].Password {
		err := errors.New("Username or password not valid")
		return err
	}
	logged = username
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
