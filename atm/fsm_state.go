package atm

import ()

// states the atm finite state machine can be in

const (
	// atm is starting up
	startingState = "starting"
	// atm is ready for requests
	readyState = "ready"
	// atm is prompting the user for their pan
	panPromptState = "pan_prompt"
	// the fsm package does not allow re-advancing to the current state so we use a loop state
	panPromptLoopState = "pan_prompt_loop"
	// atm is checking with the banking api that the account number is valid
	panCheckState = "pan_check"
	// atm admin alert that the admin menu is ready for requests
	adminMenuReadyState = "admin_menu_ready"
	// atm menu
	adminMenuState = "admin_menu"
	// loop state to redirect from the admin menu back to the admin menu
	adminMenuLoopState = "admin_menu_loop"
	// atm is prompting for a withdrawl amount
	withdrawlPromptState = "withdrawl_prompt"
	// loop state to redirect from the withdrawl prompt back to the withdrawl prompt
	withdrawlPromptLoopState = "withdrawl_prompt_loop"
	// atm is validating the withdrawl
	withdrawlState = "withdrawl"
	// atm is dispensing the withdrawl
	dispensingState = "dispensing"
	// atm transaction has finished
	endingState = "ending"
)
