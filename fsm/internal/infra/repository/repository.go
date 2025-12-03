package repository

import "github.com/mateusrangel/kit/fsm/internal/domain"

type DisputeRepotistory interface {
	CreateDispute(d *domain.Dispute) error
	UpdateState(id, newState string) error
}
