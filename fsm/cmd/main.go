package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mateusrangel/kit-examples/fsm/internal/application/service"
	"github.com/mateusrangel/kit-examples/fsm/internal/domain"
	"github.com/mateusrangel/kit-examples/fsm/internal/infra/repository"
	"github.com/mateusrangel/kit/fsm"
)

func main() {
	dispute := domain.NewDispute(uuid.New().String(), "RECEIVED")

	repo, err := repository.New()
	if err != nil {
		panic(err)
	}

	err = repo.CreateDispute(dispute)
	if err != nil {
		panic(err)
	}

	disputeService := service.NewDisputeService(dispute, repo)

	fmt.Printf("BEFORE: %v\n", disputeService.FSM.Current())
	_ = disputeService.FSM.ExecEvent("VALIDATION_SUCCEEDED")
	fmt.Printf("AFTER: %v\n", disputeService.FSM.Current())

	fmt.Printf("BEFORE: %v\n", disputeService.FSM.Current())
	_ = disputeService.FSM.ExecEvent("CLAIM_CREATED")
	fmt.Printf("AFTER: %v\n", disputeService.FSM.Current())

	fmt.Println(fsm.Visualize(disputeService.FSM))
}
