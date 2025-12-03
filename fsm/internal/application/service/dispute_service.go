package service

import (
	"fmt"

	"github.com/mateusrangel/kit/fsm"
	"github.com/mateusrangel/kit/fsm/internal/domain"
	"github.com/mateusrangel/kit/fsm/internal/infra/repository"
)

// States
const (
	StateReceived    = "RECEIVED"
	StateCreateClaim = "CREATE_CLAIM"
	StateProcessing  = "PROCESSING"
	StateFinished    = "FINISHED"
)

// Events
const (
	EventValidationSucceded = "VALIDATION_SUCCEEDED"
	EventValidationFailed   = "VALIDATION_FAILED"
	EventClaimCreated       = "CLAIM_CREATED"
	EventDisputeWon         = "DISPUTE_WON"
	EventDispotLost         = "DISPUTE_LOST"
)

type DisputeService struct {
	Dispute *domain.Dispute
	Repo    repository.DisputeRepotistory
	FSM     *fsm.FSM
}

func (o *DisputeService) UpdateState() bool {
	err := o.Repo.UpdateState(o.Dispute.Id, o.FSM.Current())
	return err != nil
}

func (o *DisputeService) SendWarningMail() bool {
	fmt.Printf("EMAIL: DISPUTE %s STATE WAS TRANSITIONED TO %s\n", o.Dispute.Id, o.FSM.Current())
	return true
}

func NewDisputeService(d *domain.Dispute, r repository.DisputeRepotistory) *DisputeService {
	disputeService := &DisputeService{Dispute: d, Repo: r}
	var states = []string{StateReceived, StateCreateClaim, StateProcessing, StateFinished}
	var events = []string{EventValidationSucceded, EventValidationFailed, EventClaimCreated, EventDisputeWon, EventDispotLost}
	transitions := []*fsm.Transition{
		{Event: EventValidationSucceded, Src: StateReceived, Dst: StateCreateClaim, Actions: []fsm.Action{disputeService.UpdateState}},
		{Event: EventValidationFailed, Src: StateReceived, Dst: StateFinished, Actions: []fsm.Action{disputeService.UpdateState, disputeService.SendWarningMail}},
		{Event: EventClaimCreated, Src: StateCreateClaim, Dst: StateProcessing, Actions: []fsm.Action{disputeService.UpdateState}},
		{Event: EventDisputeWon, Src: StateProcessing, Dst: StateFinished, Actions: []fsm.Action{disputeService.UpdateState}},
		{Event: EventDispotLost, Src: StateProcessing, Dst: StateFinished, Actions: []fsm.Action{disputeService.UpdateState, disputeService.SendWarningMail}},
	}
	m, err := fsm.New(d.State, states, events, transitions)
	if err != nil {
		panic(err)
	}
	disputeService.FSM = m
	return disputeService
}
