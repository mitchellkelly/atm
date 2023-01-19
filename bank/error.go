package bank

import (
	"fmt"
)

var InvalidUserError = fmt.Errorf("No user exists with that account number")
