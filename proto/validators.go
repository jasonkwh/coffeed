package proto

import (
	"github.com/jasonkwh/coffeed/internal/messages"
	"github.com/jasonkwh/gatekeeper"
)

func RegisterValidators() {
	gatekeeper.RegisterRequests(
		&SetBusyRequest{},
	)
}

func (r *SetBusyRequest) Validate() error {
	if r.TimeSpec == nil {
		return messages.ErrTimeSpecRequired
	}

	return nil
}
