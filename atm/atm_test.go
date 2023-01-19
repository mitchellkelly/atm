package atm

import (
	"testing"
)

var atm ATM

func init() {
	atm = ATM{
		SingleWithdrawlLimit:     500,
		DailyWithdrawlLimit:      1000,
		DailyWithdrawlCountLimit: 3,
	}
}

// test user attempting to withdraw more than their daily limit returns an error
func TestWithdrawlSingleWithdrawlLimitValidError(t *testing.T) {
	atm.Initialize()

	var err = atm.Withdraw("1", atm.SingleWithdrawlLimit+1)
	if err == nil {
		t.Errorf("An expected error did not occur.")
	}
}

// test user attempting to withdraw more than their daily limit returns an error
func TestWithdrawlDailyWithdrawlLimitValidError(t *testing.T) {
	atm.Initialize()

	var err = atm.Withdraw("1", 500)
	if err != nil {
		t.Errorf("An unxepected error occured: %s.", err)
	}
	err = atm.Withdraw("1", 500)
	if err != nil {
		t.Errorf("An unxepected error occured: %s.", err)
	}
	// now we should have reached our limit
	err = atm.Withdraw("1", 500)
	if err == nil {
		t.Errorf("An expected error did not occur.")
	}
}

// test user attempting to withdraw more than their daily limit returns an error
func TestWithdrawlDailyWithdrawlCountLimitValidError(t *testing.T) {
	atm.Initialize()

	var err = atm.Withdraw("1", 1)
	if err != nil {
		t.Errorf("An unxepected error occured: %s.", err)
	}
	err = atm.Withdraw("1", 1)
	if err != nil {
		t.Errorf("An unxepected error occured: %s.", err)
	}
	err = atm.Withdraw("1", 1)
	if err != nil {
		t.Errorf("An unxepected error occured: %s.", err)
	}

	// now we should have reached our limit
	err = atm.Withdraw("1", 1)
	if err == nil {
		t.Errorf("An expected error did not occur.")
	}
}

// test user attempting to withdraw more than their daily limit returns an error
func TestWithdrawlATMInvalidFundsValidError(t *testing.T) {
	atm.Initialize()

	var atmIntialBalance = atm.Balance
	// set the atm balance below the SingleWithdrawlLimit
	atm.Balance = atm.SingleWithdrawlLimit - 1

	var err = atm.Withdraw("1", atm.SingleWithdrawlLimit)
	atm.Balance = atmIntialBalance
	if err == nil {
		t.Errorf("An expected error did not occur.")
	}
}
