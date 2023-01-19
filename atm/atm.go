package atm

import (
	"context"
	"fmt"
	"time"

	"github.com/looplab/fsm"
	"github.com/mitchellkelly/atm/bank"
)

// ATM contains the fields and logic for making operations on a physical atm
type ATM struct {
	stopped                  bool
	bankClient               *bank.Client
	Balance                  float32
	SingleWithdrawlLimit     float32
	DailyWithdrawlLimit      float32
	DailyWithdrawlCountLimit uint
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
	// TODO read the bank client token and create an authorized client
	self.bankClient = bank.NewClient()

	// update the Balance field with the amount of money in the machine
	self.UpdateBalance()
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
// NOTE: this function will block until the atm is shutdown
func (self *ATM) Run() { // TODO logging
	// make sure the atm is initialized
	self.Initialize()

	go func() {
		// check the internal balance of the machine every day at midnight and reset the daily withdraw limits
		for self.stopped == false {
			// get the amount of time until midnight so we can set a timer
			var waitInterval = midnight().Sub(time.Now())
			// block until midnight
			<-time.After(waitInterval)

			// update balance
			self.UpdateBalance()
		}
	}()

	// create the fsm that runs the atm logic
	var atmFsm = fsm.NewFSM(startingState, atmFsmEvents, atmFsmCallbacks)
	// store the atm in the fsm metadata so the fsm can make operations on it
	atmFsm.SetMetadata(fsmAtmMetadataKey, self)
	// start the fsm
	atmFsm.Event(context.Background(), startingSuccessEvent)
}

// get all users that can access the atm
func (self ATM) UsersContext(ctx context.Context) ([]bank.User, error) {
	// get all users from the bank api
	var users, err = self.bankClient.GetUsers()

	// TODO modify users list based on atm logic (i.e. if they are allowed to access this atm)

	return users, err
}

// same as UsersContext but with a background context
func (self ATM) Users() ([]bank.User, error) {
	return self.UsersContext(context.Background())
}

// get a specific user by account number
func (self ATM) UserContext(ctx context.Context, accountNumber string) (bank.User, error) {
	// get a user from the bank api
	var user, err = self.bankClient.GetUser(accountNumber)

	// TODO modify users list based on atm logic (i.e. if they are allowed to access this atm)

	return user, err
}

// same as UserContext but with a background context
func (self ATM) User(accountNumber string) (bank.User, error) {
	return self.UserContext(context.Background(), accountNumber)
}

func (self ATM) validateWithdrawlParams(user bank.User, amount float32) error {
	// TODO check the user has enough money in their account for the withdrawl

	// check the withdraw amount is less than the SingleWithdrawlLimit
	if amount > self.SingleWithdrawlLimit {
		return fmt.Errorf("The requested withdrawl amount is greater than this ATM allows. Please request a withdrawl less than %f.",
			self.SingleWithdrawlLimit)
	}
	var remainingWithdrawlLimit = self.DailyWithdrawlLimit - user.PeriodWithdrawlSum
	// check the withdraw amount + previous withdraws is less than a user's DailyWithdrawlLimit
	if amount > remainingWithdrawlLimit {
		return fmt.Errorf("The requested withdrawl amount is greater than your allowed daily withdrawl limit. You have %f remaining in your daily withdrawl limit",
			remainingWithdrawlLimit)
	}
	// check the ATM has enough money to process the withdrawl
	if amount > self.Balance {
		return fmt.Errorf("This ATM does not have enough money to process your transaction. Please try again tomorrow.")
	}

	return nil
}

// withdraw money from a user's account
func (self *ATM) WithdrawContext(ctx context.Context, accountNumber string, amount float32) error {
	// get user info
	var user, err = self.User(accountNumber)
	if err != nil {
		return err
	}

	err = self.validateWithdrawlParams(user, amount)
	if err != nil {
		return err
	}

	// update the user's account
	err = self.bankClient.Withdraw(accountNumber, amount)
	if err != nil {
		return err
	}

	// decrement the withdrawl amount from the atm balance
	self.Balance = self.Balance - amount

	return err
}

// same as WithdrawContext but with a background context
func (self *ATM) Withdraw(accountNumber string, amount float32) error {
	return self.WithdrawContext(context.Background(), accountNumber, amount)
}

// deposit money into a user's account
func (self *ATM) DepositContext(ctx context.Context, amount float32) error {
	// TODO future release

	return fmt.Errorf("Unsupported function")
}

// same as DepositContext but with a background context
func (self *ATM) Deposit(amount float32) error {
	return self.DepositContext(context.Background(), amount)
}

// shutdown the atm
func (self *ATM) Shutdown() error {
	return nil // TODO gracefully shutdown the atm
}
