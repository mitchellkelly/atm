package atm

import (
	"context"
	"fmt"
	"strconv"

	"github.com/looplab/fsm"
)

// key for storing the atm object in the fsm
var fsmAtmMetadataKey = "atm_key"

// events defining which events progress the fsm between which states
var atmFsmEvents = fsm.Events{
	{Name: startingSuccessEvent, Src: []string{startingState}, Dst: readyState},
	{Name: readySentEvent, Src: []string{readyState}, Dst: panPromptState},
	{Name: validPanInputEvent, Src: []string{panPromptState}, Dst: panCheckState},
	{Name: invalidPanInputEvent, Src: []string{panPromptState}, Dst: panPromptLoopState},
	{Name: panInputLoopEvent, Src: []string{panPromptLoopState}, Dst: panPromptState},
	{Name: exitPrompRequestEvent, Src: []string{withdrawlPromptState}, Dst: endingState},
	{Name: adminMenuRequestEvent, Src: []string{panPromptState}, Dst: adminMenuReadyState},
	{Name: adminMenuReadySentEvent, Src: []string{adminMenuReadyState, adminMenuLoopState}, Dst: adminMenuState},
	{Name: adminMenuReturnEvent, Src: []string{adminMenuState}, Dst: panPromptState},
	{Name: adminMenuLoopEvent, Src: []string{adminMenuState}, Dst: adminMenuLoopState},
	{Name: panCheckSuccessEvent, Src: []string{panCheckState}, Dst: withdrawlPromptState},
	{Name: panCheckFailEvent, Src: []string{panCheckState}, Dst: panPromptState},
	{Name: validWithdrawlInputEvent, Src: []string{withdrawlPromptState}, Dst: withdrawlState},
	{Name: invalidWithdrawlInputEvent, Src: []string{withdrawlPromptState}, Dst: withdrawlPromptLoopState},
	{Name: withdrawlInputLoopEvent, Src: []string{withdrawlPromptLoopState}, Dst: withdrawlPromptState},
	{Name: validWithdrawlEvent, Src: []string{withdrawlState}, Dst: dispensingState},
	{Name: invalidWithdrawlEvent, Src: []string{withdrawlState}, Dst: withdrawlPromptState},
	{Name: dispensingCompleteEvent, Src: []string{dispensingState}, Dst: endingState},
	{Name: endingSentEvent, Src: []string{endingState}, Dst: readyState},
}

// callbacks that happen after a state change of the fsm
// this is the meat of the atm logic
var atmFsmCallbacks = fsm.Callbacks{
	readyState: func(ctx context.Context, event *fsm.Event) {
		fmt.Println("Welcome to the Mitchell Kelly ATM!")

		event.FSM.Event(ctx, readySentEvent)
	},
	panPromptState: func(ctx context.Context, event *fsm.Event) {
		fmt.Print("Please enter your Personal Access Number or 'a' for the admin menu: ")

		var userInput string
		// accept user input
		fmt.Scanln(&userInput)

		switch userInput {
		case "a":
			// admin menu requst event
			event.FSM.Event(ctx, adminMenuRequestEvent)
		default:
			var validPanFormat bool

			// TODO validate the pan format
			validPanFormat = true

			if validPanFormat {
				// valid pan input event
				event.FSM.Event(ctx, validPanInputEvent, userInput)
			} else {
				fmt.Println("An invalid Personal Access Number format was provided")
				// invalid pan input event
				event.FSM.Event(ctx, invalidPanInputEvent)
			}

		}
	},
	panPromptLoopState: func(ctx context.Context, event *fsm.Event) {
		// advance to pan input
		event.FSM.Event(ctx, panInputLoopEvent)
	},
	adminMenuReadyState: func(ctx context.Context, event *fsm.Event) {
		fmt.Println("Welcome to the super secret admin menu.")

		// advance to the admin menu prompt
		event.FSM.Event(ctx, adminMenuReadySentEvent)
	},
	adminMenuState: func(ctx context.Context, event *fsm.Event) {
		fmt.Println("Enter 'l' to list users")
		fmt.Println("Enter 'p' to return to the previous screen")
		fmt.Print("Please enter an action from the list above: ")

		var userInput string
		// accept user input
		fmt.Scanln(&userInput)

		switch userInput {
		case "l":
			// get the atm object from the fsm
			var metadata, _ = event.FSM.Metadata(fsmAtmMetadataKey)
			var atm, _ = metadata.(*ATM)

			// get all of the atm users
			var users, err = atm.UsersContext(ctx)
			if err == nil {
				// print the atm users
				fmt.Println("Users:")
				for _, x := range users {
					fmt.Println(x)
				}
			} else {
				fmt.Println("An error occured requesting data from the banking api. Please try again later.")
			}

			// loop back to the admin menu to process more requests
			event.FSM.Event(ctx, adminMenuLoopEvent)
		case "p":
			// return to the previous state
			event.FSM.Event(ctx, adminMenuReturnEvent)
		default:
			fmt.Println("Invalid input. Please try again")
			// loop back to the admin menu to allow the user to try again
			event.FSM.Event(ctx, adminMenuLoopEvent)
		}
	},
	adminMenuLoopState: func(ctx context.Context, event *fsm.Event) {
		// advance to the admin menu
		event.FSM.Event(ctx, adminMenuReadySentEvent)
	},
	panCheckState: func(ctx context.Context, event *fsm.Event) {
		fmt.Println("Checking the Personal Access Number provided, please wait.")

		// get the atm object from the fsm
		var metadata, _ = event.FSM.Metadata(fsmAtmMetadataKey)
		var atm, _ = metadata.(*ATM)

		var userPan string
		// get the user provided pan from the event args
		if len(event.Args) > 0 {
			userPan, _ = event.Args[0].(string)
		}

		// get user info using the provided pan
		var user, err = atm.UserContext(ctx, userPan)
		if err == nil {
			// bank api pan check was successful event
			event.FSM.Event(ctx, panCheckSuccessEvent, user)
		} else {
			fmt.Println(err)
			// bank api pan check was unsuccessful event
			event.FSM.Event(ctx, panCheckFailEvent, event.Args...)
		}
	},
	withdrawlPromptState: func(ctx context.Context, event *fsm.Event) {
		fmt.Print("Please enter the amount you would like to withdraw, or 'q' to quit: ")

		var userInput string
		// accept user input
		fmt.Scanln(&userInput)

		switch userInput {
		case "q":
			// leave the prompt menu
			event.FSM.Event(ctx, exitPrompRequestEvent)
		default:
			var user User
			// get the user provided pan from the event args
			if len(event.Args) > 0 {
				user, _ = event.Args[0].(User)
			}

			// parse the user input as a float
			var withdrawlAmount, err = strconv.ParseFloat(userInput, 32)
			if err == nil {
				// withdrawl user input parsing was successful event
				event.FSM.Event(ctx, validWithdrawlInputEvent, user, float32(withdrawlAmount))
			} else {
				fmt.Println("Invalid input was provided. Please input a valid number")
				// withdrawl user input parsing was unsuccessful event
				event.FSM.Event(ctx, invalidWithdrawlInputEvent, event.Args...)
			}
		}
	},
	withdrawlState: func(ctx context.Context, event *fsm.Event) {
		fmt.Println("Processing the withdrawl, please wait")

		// get the atm object from the fsm
		var metadata, _ = event.FSM.Metadata(fsmAtmMetadataKey)
		var atm, _ = metadata.(*ATM)

		var user User
		var withdrawlAmount float32
		// get the user provided pan and withdrawl amount from the event args
		if len(event.Args) >= 2 {
			user, _ = event.Args[0].(User)
			withdrawlAmount, _ = event.Args[1].(float32)
		}

		// attempt to make a withdrawl from the atm
		var err = atm.WithdrawContext(ctx, user.AccountNumber, withdrawlAmount)
		if err == nil {
			// withdrawl was successful event
			event.FSM.Event(ctx, validWithdrawlEvent)
		} else {
			fmt.Println(err)
			// withdrawl was unsuccessful event
			event.FSM.Event(ctx, invalidWithdrawlEvent, event.Args...)
		}
	},
	withdrawlPromptLoopState: func(ctx context.Context, event *fsm.Event) {
		// loop back to the withdrawl prompt
		event.FSM.Event(ctx, withdrawlInputLoopEvent, event.Args...)
	},
	dispensingState: func(ctx context.Context, event *fsm.Event) {
		fmt.Println("The withdrawl has processed successfully. The money is being dispensed now.")
		// money was dispensed event
		event.FSM.Event(ctx, dispensingCompleteEvent)
	},
	endingState: func(ctx context.Context, event *fsm.Event) {
		fmt.Println("Thanks for using the Mitchell Kelly ATM! Please come back soon.")
		fmt.Println()
		// transaction complete event
		event.FSM.Event(ctx, endingSentEvent)
	},
}
