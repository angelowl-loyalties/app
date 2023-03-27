package tests

import (
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/internal"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/models"
	"github.com/google/uuid"
)

// global variable for seed transactions for testing
var SeedTransactions = make(map[string]models.Transaction)
// var SeedBaseCampaigns = make(map[string]models.Campaign)
// var SeedPromoCampaigns = make(map[string]models.Campaign)
// var SeedExclusions = make(map[string]models.Exclusion)

func Test_IsExcluded(t *testing.T) {
	etcdAddSeedData()

	type args struct {
		transactionDate time.Time
		mcc             int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"exclusion_mcc_not_excluded", args{time.Now(), 6969}, false},
		{"exclusion_mcc_excluded", args{time.Now(), 4125}, true},
		{"exclusion_mcc_excluded_valid_tomorrow", args{time.Now(), 5001}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.IsExcluded(tt.args.transactionDate, tt.args.mcc); got != tt.want {
				t.Errorf("IsExcluded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMatchingCampaigns(t *testing.T) {
	etcdAddSeedData()

	baseMatchingCampaigns := []models.Campaign{}
	baseMatchingCampaigns = append(baseMatchingCampaigns, internal.BaseCampaignsEtcd["001"])
	baseMatchingCampaigns = append(baseMatchingCampaigns, internal.BaseCampaignsEtcd["005"])
	promoMatchingCampaigns := []models.Campaign{}
	wantCampaign_base_2_promo_0 := [][]models.Campaign{}
	wantCampaign_base_2_promo_0 = append(wantCampaign_base_2_promo_0, baseMatchingCampaigns)
	wantCampaign_base_2_promo_0 = append(wantCampaign_base_2_promo_0, promoMatchingCampaigns)

	type args struct {
		transaction models.Transaction
	}
	tests := []struct {
		name         string
		args         args
		wantCampaign [][]models.Campaign
	}{
		{"base_2_promo_0", args{SeedTransactions["001"]}, wantCampaign_base_2_promo_0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCampaign := internal.GetMatchingCampaigns(tt.args.transaction)
			gotBaseCampaign := gotCampaign[0]
			gotPromoCampaign := gotCampaign[1]

			// Sort the gotten Base and Promo Campaigns, to ensure order is consistent over different test runs
			sort.Slice(gotBaseCampaign, func(i, j int) bool { return gotBaseCampaign[i].ID.String() < gotBaseCampaign[j].ID.String() })
			sort.Slice(gotPromoCampaign, func(i, j int) bool { return gotPromoCampaign[i].ID.String() < gotPromoCampaign[j].Name })

			if !reflect.DeepEqual(gotCampaign, tt.wantCampaign) {
				t.Errorf("getMatchingCampaigns() = %v, want %v", gotCampaign, tt.wantCampaign)
			}
		})
	}
}

func Test_isCampaignMatch(t *testing.T) {
	etcdAddSeedData()

	type args struct {
		campaign    models.Campaign
		transaction models.Transaction
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// All "false" cases are only a single 1 change from the first "true" case
		{"match", args{internal.BaseCampaignsEtcd["001"], SeedTransactions["001"]}, true},
		{"no_match_card_type", args{internal.BaseCampaignsEtcd["001"], SeedTransactions["002"]}, false},
		{"no_match_campaign_not_started", args{internal.BaseCampaignsEtcd["002"], SeedTransactions["001"]}, false},
		{"no_match_campaign_ended", args{internal.BaseCampaignsEtcd["003"], SeedTransactions["001"]}, false},
		{"no_match_foreign_sgd", args{internal.BaseCampaignsEtcd["001"], SeedTransactions["003"]}, false},
		{"no_match_different_merchant", args{internal.BaseCampaignsEtcd["001"], SeedTransactions["004"]}, false},
		{"no_match_min_spend_not_met", args{internal.BaseCampaignsEtcd["001"], SeedTransactions["005"]}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.IsCampaignMatch(tt.args.campaign, tt.args.transaction); got != tt.want {
				t.Errorf("isCampaignMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_ProcessMessageJSON(t *testing.T) {
// 	messageJSON1 := `{
// 		"id":              "4aab2f7c-4dd3-4a77-beb8-8582048c9bdb",
// 		"card_id":          "3c0b3d7f-c011-4a7d-b47e-1f7c03a8ca53",
// 		"merchant":        "Best Buy",
// 		"mcc":             5912,
// 		"currency":        "USD",
// 		"amount":          500.00,
// 		"sgd_amount":       712.00,
// 		"transaction_id":   "1234abcd",
// 		"transaction_date": "2023-01-23",
// 		"card_pan":         "1234567890123456",
// 		"card_type":        "Points"
// 	}`

// 	// models.InitDB()

// 	tests := []struct {
// 		name        string
// 		messageJSON string
// 		gotErr      bool
// 	}{
// 		{"invalid_message", "invalid_json", true},
// 		{"valid_message", messageJSON1, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Test
// 			err := ProcessMessageJSON(tt.messageJSON)

// 			// Assert
// 			if (err != nil) != tt.gotErr {
// 				t.Errorf("ProcessMessageJSON() error = %v, gotErr %v", err, tt.gotErr)
// 			}
// 		})
// 	}
// }

func etcdAddSeedData() {
	// Add Exclusions
	internal.ExclusionsEtcd["001"] = models.Exclusion{
		ID:        uuid.MustParse("e38adb10-a96a-4b55-aebd-7cdc9b973e7b"),
		MCC:       4125,
		ValidFrom: time.Now(),
	}
	internal.ExclusionsEtcd["002"] = models.Exclusion{
		ID:        uuid.MustParse("e38adb10-a96a-4b55-aebd-7cdc9b973e7b"),
		MCC:       5001,
		ValidFrom: time.Now().AddDate(0, 0, 1),
	}

	// Add Base Campaigns
	// Valid campaign
	internal.BaseCampaignsEtcd["001"] = models.Campaign{
		ID:                 uuid.MustParse("7b1f04eb-f54c-4f9d-8baf-a4c00dddf111"),
		Name:               "Summer Sale",
		MinSpend:           50.0,
		Start:              time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		End:                time.Date(2023, 8, 31, 23, 59, 59, 0, time.UTC),
		RewardProgram:      "Points",
		RewardAmount:       500,
		MCC:                "7011,5963",
		Merchant:           "Best Buy",
		IsBaseReward:		true,
		ForForeignCurrency: true,
	}

	internal.BaseCampaignsEtcd["002"] = models.Campaign{
		ID:                 uuid.MustParse("1c7f6942-85f9-4a9a-b1ab-6dab27c94222"),
		Name:               "Winter Warmup",
		MinSpend:           100.0,
		Start:              time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		End:                time.Date(2024, 2, 28, 23, 59, 59, 0, time.UTC),
		RewardProgram:      "Cashback",
		RewardAmount:       25,
		MCC:                "5913",
		Merchant:           "Grab",
		IsBaseReward:		true,
		ForForeignCurrency: true,
	}

	internal.BaseCampaignsEtcd["003"] = models.Campaign{
		ID:                 uuid.MustParse("1c7f6942-85f9-4a9a-b1ab-6dab27c94333"),
		Name:               "Winter Warmup",
		MinSpend:           100.0,
		Start:              time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC),
		End:                time.Date(2023, 2, 28, 23, 59, 59, 0, time.UTC),
		RewardProgram:      "Cashback",
		RewardAmount:       25,
		MCC:                "5913",
		Merchant:           "Grab",
		IsBaseReward:		false,
		ForForeignCurrency: false,
	}

	internal.BaseCampaignsEtcd["004"] = models.Campaign{
		ID:                 uuid.MustParse("ddb0a58f-6dca-41f3-a3a9-d40961670444"),
		Name:               "Spring Fling",
		MinSpend:           0.0,
		Start:              time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		End:                time.Date(2024, 5, 31, 23, 59, 59, 0, time.UTC),
		RewardProgram:      "Visa",
		RewardAmount:       300,
		MCC:                "5963",
		Merchant:           "",
		IsBaseReward:		false,
		ForForeignCurrency: false,
	}

	internal.BaseCampaignsEtcd["005"] = models.Campaign{
		ID:                 uuid.MustParse("ddb0a58f-6dca-41f3-a3a9-d40961670555"),
		Name:               "Summer x2",
		MinSpend:           0.0,
		Start:              time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		End:                time.Date(2024, 5, 31, 23, 59, 59, 0, time.UTC),
		RewardProgram:      "Points",
		RewardAmount:       600,
		MCC:                "5963,7011",
		Merchant:           "",
		IsBaseReward:		true,
		ForForeignCurrency: true,
		
	}

	// Add Transactions
	SeedTransactions["001"] = models.Transaction{
		ID:              uuid.MustParse("4aab2f7c-4dd3-4a77-beb8-8582048c9bdb"),
		CardID:          uuid.MustParse("3c0b3d7f-c011-4a7d-b47e-1f7c03a8ca53"),
		Merchant:        "Best Buy",
		MCC:             7011,
		Currency:        "USD",
		Amount:          500.00,
		SGDAmount:       712.00,
		TransactionID:   "1234abcd",
		TransactionDate: "2023-10-23",
		CardPAN:         "1234567890123456",
		CardType:        "Points",
	}

	SeedTransactions["002"] = models.Transaction{
		ID:              uuid.MustParse("4aab2f7c-4dd3-4a77-beb8-8582048c9bdb"),
		CardID:          uuid.MustParse("3c0b3d7f-c011-4a7d-b47e-1f7c03a8ca53"),
		Merchant:        "Best Buy",
		MCC:             5912,
		Currency:        "USD",
		Amount:          500.00,
		SGDAmount:       712.00,
		TransactionID:   "1234abcd",
		TransactionDate: "2023-10-23",
		CardPAN:         "1234567890123456",
		CardType:        "Visa",
	}

	SeedTransactions["003"] = models.Transaction{
		ID:              uuid.MustParse("4aab2f7c-4dd3-4a77-beb8-8582048c9bdb"),
		CardID:          uuid.MustParse("3c0b3d7f-c011-4a7d-b47e-1f7c03a8ca53"),
		Merchant:        "Best Buy",
		MCC:             5912,
		Currency:        "SGD",
		SGDAmount:       712.00,
		TransactionID:   "1234abcd",
		TransactionDate: "2023-10-23",
		CardPAN:         "1234567890123456",
		CardType:        "Points",
	}

	SeedTransactions["004"] = models.Transaction{
		ID:              uuid.MustParse("4aab2f7c-4dd3-4a77-beb8-8582048c9bdb"),
		CardID:          uuid.MustParse("3c0b3d7f-c011-4a7d-b47e-1f7c03a8ca53"),
		Merchant:        "Walgreens",
		MCC:             5912,
		Currency:        "USD",
		Amount:          500.00,
		SGDAmount:       712.00,
		TransactionID:   "1234abcd",
		TransactionDate: "2023-10-23",
		CardPAN:         "1234567890123456",
		CardType:        "Points",
	}

	SeedTransactions["005"] = models.Transaction{
		ID:              uuid.MustParse("4aab2f7c-4dd3-4a77-beb8-8582048c9bdb"),
		CardID:          uuid.MustParse("3c0b3d7f-c011-4a7d-b47e-1f7c03a8ca53"),
		Merchant:        "Best Buy",
		MCC:             5912,
		Currency:        "USD",
		Amount:          10.00,
		SGDAmount:       14.24,
		TransactionID:   "1234abcd",
		TransactionDate: "2023-10-23",
		CardPAN:         "1234567890123456",
		CardType:        "Points",
	}
}
