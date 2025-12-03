package domain

type Dispute struct {
	Id    string
	State string
}

func NewDispute(id string, state string) *Dispute {
	return &Dispute{Id: id, State: state}
}
