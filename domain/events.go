package domain

import (
	"github.com/eduncan911/es"
)

func i(contract string, eventID EventID) *es.Info {
	return es.NewInfo(contract, eventID.ID)
}
