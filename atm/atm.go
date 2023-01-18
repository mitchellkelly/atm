package atm

import (
	"context"
	"fmt"
	"time"

	"github.com/mitchellkelly/atm/bank"
)

// ATM contains the fields and logic for making operations on a physical atm
type ATM struct {
	stopped              bool
	bankClient           *bank.Client
	Balance              float32
	SingleWithdrawlLimit float32
	DailyWithdrawlLimit  float32
	Withdrawls           map[string]float32
}

// count the amount of money inside of the machine
func (self ATM) InternalBalance() float32 {
	return 5000 // TODO count the money inside of the machine
}

// set the Balance field to the amount of money inside of the machine
func (self *ATM) UpdateBalance() {
	self.Balance = self.InternalBalance()
}

// initialize the atm with some values
func (self *ATM) Initialize() {
	// TODO
}

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

// set up the atm to an operational state and run the the finite state machine
func (self *ATM) Run() { // TODO logging
	// TODO
}

// get all users that can access the atm
func (self ATM) UsersContext(ctx context.Context) ([]User, error) {
	var users = make([]User, 0)
	var err error

	// TODO

	return users, err
}

// same as UsersContext but with a background context
func (self ATM) Users() ([]User, error) {
	return self.UsersContext(context.Background())
}

// get a specific user by account number
func (self ATM) UserContext(ctx context.Context, accountNumber string) (User, error) {
	var user User
	var err error

	// TODO

	return user, err
}

// same as UserContext but with a background context
func (self ATM) User(accountNumber string) (User, error) {
	return self.UserContext(context.Background(), accountNumber)
}

func (self ATM) validateWithdrawlParams(user User, amount float32) error {
	var err error

	// TODO

	return err
}

// withdraw money from a user's account
func (self *ATM) WithdrawContext(ctx context.Context, accountNumber string, amount float32) error {
	var err error

	// TODO

	return err
}

// same as WithdrawContext but with a background context
func (self *ATM) Withdraw(accountNumber string, amount float32) error {
	return self.WithdrawContext(context.Background(), accountNumber, amount)
}

// deposit money into a user's account
func (self *ATM) DepositContext(ctx context.Context, amount float32) error {
	var err error

	// TODO future release
	err = fmt.Errorf("Unsupported function")

	return err
}

// same as DepositContext but with a background context
func (self *ATM) Deposit(amount float32) error {
	return self.DepositContext(context.Background(), amount)
}

// shutdown the atm
func (self *ATM) Shutdown() error {
	return nil // TODO gracefully shutdown the atm
}
