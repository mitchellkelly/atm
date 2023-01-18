package atm

import (
	"context"
	"fmt"
	"strconv"

	"github.com/looplab/fsm"
)

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
		// TODO
	},
	panPromptState: func(ctx context.Context, event *fsm.Event) {
		// TODO
	},
	panPromptLoopState: func(ctx context.Context, event *fsm.Event) {
		// TODO
	},
	adminMenuReadyState: func(ctx context.Context, event *fsm.Event) {
		// TODO
	},
	adminMenuState: func(ctx context.Context, event *fsm.Event) {
		// TODO
	},
	adminMenuLoopState: func(ctx context.Context, event *fsm.Event) {
		// TODO
	},
	panCheckState: func(ctx context.Context, event *fsm.Event) {
		// TODO
	},
	withdrawlPromptState: func(ctx context.Context, event *fsm.Event) {
		// TODO
	},
	withdrawlState: func(ctx context.Context, event *fsm.Event) {
		// TODO
	},
	withdrawlPromptLoopState: func(ctx context.Context, event *fsm.Event) {
		// TODO
	},
	dispensingState: func(ctx context.Context, event *fsm.Event) {
		// TODO
	},
	endingState: func(ctx context.Context, event *fsm.Event) {
		// TODO
	},
}
