package repository

import "github.com/mateusrangel/kit-examples/fsm/internal/domain"

type DisputeRepotistory interface {
	CreateDispute(d *domain.Dispute) error
	UpdateState(id, newState string) error
}
