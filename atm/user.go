package atm

import (
	"fmt"

	"github.com/mitchellkelly/atm/bank"
)

type User struct {
	bank.User
	DailyWithdrawls float32
}

func (self User) String() string {
	return fmt.Sprintf("Name: %s, Account Number: %s", self.User.Name, self.User.AccountNumber)
}
