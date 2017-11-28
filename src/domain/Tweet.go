package domain

import "time"

type Tweet struct {
	Username string
	Text     string
	Date     *time.Time
}

func NewTweet(u string, t string) *Tweet {
	var date = time.Now()
	return &Tweet{u, t, &date}
}
