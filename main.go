package main

import (
	"github.com/mitchellkelly/atm/atm"
)

// program implmenting the behavior of an ATM

func main() {
	// intialize the atm
	var machine = atm.ATM{
		SingleWithdrawlLimit: 500,
		DailyWithdrawlLimit:  1000,
	}
	// run the atm
	machine.Run()
}
