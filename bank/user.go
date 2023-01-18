package bank

import (
	"context"
	"fmt"
)

type User struct {
	AccountNumber string
	Name          string
	Balance       float32
}

func (self Client) GetUsersContext(ctx context.Context) ([]User, error) {
	// TODO user info should be requested from an api
	var users = []User{
		User{
			AccountNumber: "1",
			Name:          "Mitchell Kelly",
			Balance:       -1,
		},
		User{
			AccountNumber: "2",
			Name:          "Antall Fernandes",
			Balance:       -1,
		},
	}

	return users, nil
}

func (self Client) GetUsers() ([]User, error) {
	return self.GetUsersContext(context.Background())
}

func (self Client) GetUserContext(ctx context.Context, accountNumber string) (User, error) {
	var user User
	var err error

	// TODO this api should request one user from an api instead of requesting all and iterating
	var users, _ = self.GetUsersContext(ctx)
	for _, x := range users {
		if accountNumber == x.AccountNumber {
			user = x
		}
	}

	// TODO this error would have came back from the api so we shouldnt need this block
	if len(user.Name) == 0 {
		err = fmt.Errorf("No user exists with that account number")
	}

	return user, err
}

func (self Client) GetUser(accountNumber string) (User, error) {
	return self.GetUserContext(context.Background(), accountNumber)
}
