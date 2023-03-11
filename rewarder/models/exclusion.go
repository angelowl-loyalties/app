package models

import (
	"time"

	"github.com/google/uuid"
)

type Exclusion struct {
	ID        uuid.UUID `json:"id"`
	MCC       int       `json:"mcc"`
	ValidFrom time.Time `json:"valid_from"`
}
