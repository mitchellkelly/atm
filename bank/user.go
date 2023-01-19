package bank

import (
	"context"
)

// canned user data
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
	User{
		AccountNumber: "3",
		Name:          "John Doe",
		Balance:       -1,
	},
	User{
		AccountNumber: "4",
		Name:          "Jane Doe",
		Balance:       -1,
	},
	User{
		AccountNumber: "5",
		Name:          "Alice",
		Balance:       -1,
	},
	User{
		AccountNumber: "6",
		Name:          "Bob",
		Balance:       -1,
	},
}

type User struct {
	AccountNumber        string
	Name                 string
	Balance              float32
	PeriodWithdrawlCount uint
	PeriodWithdrawlSum   float32
}

func (self Client) GetUsersContext(ctx context.Context) ([]User, error) {
	return self.users, nil
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
		err = InvalidUserError
	}

	return user, err
}

func (self Client) GetUser(accountNumber string) (User, error) {
	return self.GetUserContext(context.Background(), accountNumber)
}

func (self *Client) WithdrawContext(ctx context.Context, accountNumber string, amount float32) error {
	var err error

	// TODO check the user has enough money in their account for the withdrawl

	var userIndex = -1
	// find the user in the users map so we can update it
	for i, x := range self.users {
		if accountNumber == x.AccountNumber {
			userIndex = i
		}
	}

	if userIndex != -1 {
		// increment the withdrawl count
		self.users[userIndex].PeriodWithdrawlCount = self.users[userIndex].PeriodWithdrawlCount + 1
		// add the withdrawl amount to the period withdrawl sum
		self.users[userIndex].PeriodWithdrawlSum = self.users[userIndex].PeriodWithdrawlSum + amount
	} else {
		err = InvalidUserError
	}

	return err
}

func (self *Client) Withdraw(accountNumber string, amount float32) error {
	return self.WithdrawContext(context.Background(), accountNumber, amount)
}
