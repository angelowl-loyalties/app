package models

import (
	"time"

	"github.com/google/uuid"
)

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
