package atm

import ()

// events that move the atm finite state machine between states

const (
	// atm has successfully finished starting
	startingSuccessEvent = "starting_success"
	// atm has failed to start
	startingFailEvent = "starting_fail"
	// atm ready confirmation has been sent
	readySentEvent = "ready_sent"
	// user personal access number is valid
	validPanInputEvent = "valid_pan_input"
	// user personal access number is invalid
	invalidPanInputEvent = "invalid_pan_input"
	// looping back to pan input
	panInputLoopEvent = "pan_input_loop_requested"
	// admin menu requested
	adminMenuRequestEvent = "admin_menu_request"
	// admin menu ready message has been sent
	adminMenuReadySentEvent = "admin_menu_ready_sent"
	// admin menu exit requested
	adminMenuReturnEvent = "admin_menu_return"
	// looping back to admin menu input
	adminMenuLoopEvent = "admin_menu_loop_requested"
	// atm checked the bank api for the pan and it was valid
	panCheckSuccessEvent = "pan_check_success"
	// atm checked the bank api for the pan and it was invalid
	panCheckFailEvent = "pan_check_fail"
	// user supplied withdrawl input was valid
	validWithdrawlInputEvent = "valid_withdrawl_input"
	// user supplied withdrawl input was invalid
	invalidWithdrawlInputEvent = "invalid_withdrawl_input"
	// looping back to withdrawl input
	withdrawlInputLoopEvent = "withdrawl_input_loop_requested"
	// exiting a prompt
	exitPrompRequestEvent = "exit_prompt_request"
	// withdrawl request was successful
	validWithdrawlEvent = "valid_withdrawl"
	// withdrawl request was unsuccessful
	invalidWithdrawlEvent = "invalid_withdrawl"
	// withdrawl has been dispensed
	dispensingCompleteEvent = "dispensing_complete"
	// ending transaction
	endingSentEvent = "ending_sent"
)
