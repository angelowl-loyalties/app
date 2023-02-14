package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              uuid.UUID
	CardID          uuid.UUID
	Merchant        string
	MCC             int
	Currency        string
	Amount          float64
	SGDAmount       float64
	TransactionID   string
	TransactionDate string
	CardPAN         string
	CardType        string
}

type Campaign struct {
	ID            uuid.UUID
	Name          string
	MinSpend      float64
	Start         time.Time
	End           time.Time
	RewardProgram string
	RewardAmount  int
	MCC           int
	Merchant      string
	CardType      string
}

type Exclusion struct {
	ID        uuid.UUID
	MCC       int
	ValidFrom time.Time
}

var Seed_transaction = Transaction{
	ID:              uuid.MustParse("4aab2f7c-4dd3-4a77-beb8-8582048c9bdb"),
	CardID:          uuid.MustParse("3c0b3d7f-c011-4a7d-b47e-1f7c03a8ca53"),
	Merchant:        "Acme Inc.",
	MCC:             5912,
	Currency:        "SGD",
	Amount:          100.00,
	SGDAmount:       142.40,
	TransactionID:   "1234abcd",
	TransactionDate: "10-02-23",
	CardPAN:         "1234567890123456",
	CardType:        "Visa",
}

// Remove after consuming from etcd
var Seed_campaigns = []Campaign{
	{
		ID:            uuid.MustParse("7b1f04eb-f54c-4f9d-8baf-a4c00dddf3b3"),
		Name:          "Summer Sale",
		MinSpend:      50.0,
		Start:         time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		End:           time.Date(2023, 8, 31, 23, 59, 59, 0, time.UTC),
		RewardProgram: "Points",
		RewardAmount:  500,
		MCC:           7011,
		Merchant:      "Best Buy",
		CardType:      "Visa",
	},
	{
		ID:            uuid.MustParse("1c7f6942-85f9-4a9a-b1ab-6dab27c94f83"),
		Name:          "Winter Warmup",
		MinSpend:      100.0,
		Start:         time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		End:           time.Date(2024, 2, 28, 23, 59, 59, 0, time.UTC),
		RewardProgram: "Cashback",
		RewardAmount:  25,
		MCC:           5913,
		Merchant:      "Home Depot",
		CardType:      "Master",
	},
	{
		ID:            uuid.MustParse("ddb0a58f-6dca-41f3-a3a9-d40961670b5b"),
		Name:          "Spring Fling",
		MinSpend:      75.0,
		Start:         time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
		End:           time.Date(2023, 5, 31, 23, 59, 59, 0, time.UTC),
		RewardProgram: "Discount",
		RewardAmount:  10,
		MCC:           5963,
		Merchant:      "Petco",
		CardType:      "Crypto",
	},
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
