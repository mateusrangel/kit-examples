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
	repo, err := repository.New()
	if err != nil {
		panic(err)
	}

	dispute := domain.NewDispute(uuid.New().String(), "RECEIVED")
	err = repo.CreateDispute(dispute)
	if err != nil {
		panic(err)
	}

	orderFSM := service.NewDisputeService(dispute, repo)

	fmt.Printf("BEFORE: %v\n", orderFSM.FSM.Current())
	_ = orderFSM.FSM.ExecEvent("VALIDATION_SUCCEEDED")
	fmt.Printf("AFTER: %v\n", orderFSM.FSM.Current())

	fmt.Printf("BEFORE: %v\n", orderFSM.FSM.Current())
	_ = orderFSM.FSM.ExecEvent("CLAIM_CREATED")
	fmt.Printf("AFTER: %v\n", orderFSM.FSM.Current())

	fmt.Println(fsm.Visualize(orderFSM.FSM))
}
