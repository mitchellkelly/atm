package atm

import (
	"fmt"
)

var unsupportedFunctionError = fmt.Errorf("Unsupported function")

type WithdrawError interface {
	Error() string
	// no-op function WithdrawError types need to implement to differentiate them from other errors
	wError()
}

type WithdrawErrorSingleLimitMaxExceeded struct {
	RemainingLimit float32
}

func (self WithdrawErrorSingleLimitMaxExceeded) Error() string {
	return fmt.Sprintf("The requested withdrawl amount is greater than this ATM allows. Please request a withdrawl less than %f.",
		self.RemainingLimit)
}

func (self WithdrawErrorSingleLimitMaxExceeded) wError() {}

type WithdrawErrorDailyLimitMaxExceeded struct {
	RemainingLimit float32
}

func (self WithdrawErrorDailyLimitMaxExceeded) Error() string {
	return fmt.Sprintf("The requested withdrawl amount is greater than your allowed daily withdrawl limit. You have %f remaining in your daily withdrawl limit",
		self.RemainingLimit)
}

func (self WithdrawErrorDailyLimitMaxExceeded) wError() {}

type WithdrawErrorWithdrawlCountExceeded struct{}

func (self WithdrawErrorWithdrawlCountExceeded) Error() string {
	return "You have exceeded the amout of withdrawls this machine allows within a single day. Please come back tomorrow."
}

func (self WithdrawErrorWithdrawlCountExceeded) wError() {}

type WithdrawErrorATMInvalidFunds struct{}

func (self WithdrawErrorATMInvalidFunds) Error() string {
	return "This ATM does not have enough money to process your transaction. Please try again tomorrow."
}

func (self WithdrawErrorATMInvalidFunds) wError() {}
