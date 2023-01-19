package bank

import (
	"time"
)

// TOOD the client just returns canned data
// the client should connect to an api and use real data

type Client struct {
	users []User
}

// TODO taken from atm package. the api should handle updating users
func midnight() time.Time {
	// get current time
	var currentTime = time.Now()
	// add 24 hours to today to get tomorrows date
	var tomorrow = currentTime.AddDate(0, 0, 1)
	// we need midnight so we need to create a new date without hours / minutes / secs / nsecs
	var midnight = time.Date(tomorrow.Year(), tomorrow.Month(),
		tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location())

	return midnight
}

func init() {
}

func NewClient() *Client {
	var client = &Client{}
	// update the client with our canned user data
	client.updateUsers()

	go func() {
		for true {
			// get the amount of time until midnight so we can set a timer
			var waitInterval = midnight().Sub(time.Now())
			// block until midnight
			<-time.After(waitInterval)

			// clear user period values // TODO the api should handle updating users
			client.updateUsers()
		}
	}()

	return client
}

func (self *Client) updateUsers() {
	self.users = make([]User, len(users))
	copy(self.users, users)
}
