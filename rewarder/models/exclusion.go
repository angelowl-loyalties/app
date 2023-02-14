package models

import (
	"time"

	"github.com/google/uuid"
)

type Exclusion struct {
	ID        uuid.UUID
	MCC       int
	ValidFrom time.Time
}

var Seed_exclusions = []Exclusion{
	{
		ID:        uuid.MustParse("e5e2a5c5-f6e0-48e7-9ccb-7c6cfa78a873"),
		MCC:       5915,
		ValidFrom: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        uuid.MustParse("6ddb0ecc-0f05-44c2-b6c3-7587d5a56bab"),
		MCC:       7011,
		ValidFrom: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        uuid.MustParse("1e5b5fd5-69d5-4f98-8f33-1b9069d31a5b"),
		MCC:       5963,
		ValidFrom: time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        uuid.MustParse("e38adb10-a96a-4b55-aebd-7cdc9b973e7b"),
		MCC:       5977,
		ValidFrom: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
	},
}
