package bank

import ()

// TOOD the client just returns canned data
// the client should connect to an api and use real data

type Client struct{}

func NewClient() *Client {
	return &Client{}
}
