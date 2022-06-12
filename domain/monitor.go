package domain

import (
	"fmt"
	"time"
)

type Monitor struct {
	ID        string        `json:"id"`
	CreatedAt time.Time     `json:"createdAt" validate:"required"`
	Period    time.Duration `json:"period" validate:"required"`
	Frequency time.Duration `json:"frequency" validate:"required"`
	Rates     []Rate        `json:"rate" validate:"required"`
	Complete  bool          `json:"complete" validate:"required"`
}

func (m Monitor) Validate() error {
	//TODO
	if m.ID == "" {
		return fmt.Errorf("Empty monitor ID")
	}
	if m.CreatedAt.IsZero() {
		return fmt.Errorf("Empty dateCreated for monitor name %v", m.ID)
	}
	if m.Period.Seconds() < 1 {
		return fmt.Errorf("Too small period for  monitor name %v", m.ID)
	}
	if m.Frequency.Seconds() < 1 {
		return fmt.Errorf("Too small frequency for monitor name %v", m.ID)
	}

	return nil
}
