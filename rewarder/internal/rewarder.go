package internal

import (
	"encoding/json"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/models"
	"github.com/gocql/gocql"
)

const (
	layout = "02-01-06"
)

func ProcessMessageJSON(messageJSON string) error {
	var transaction models.Transaction

	json.Unmarshal([]byte(messageJSON), &transaction) // Convert JSON message to transaction object

	return ProcessMessage(transaction)
}

func ProcessMessage(transaction models.Transaction) error {
	transactionDate, _ := time.Parse(layout, transaction.TransactionDate)
	if isExcluded(transactionDate, transaction.MCC) {
		// Delta and remarks for the exclusion case
		var delta float64 = 0
		var remarks = "Campaigns don't apply"
		reward := models.Reward{
			ID:              gocql.UUID(transaction.ID),
			CardID:          gocql.UUID(transaction.CardID),
			Merchant:        transaction.Merchant,
			MCC:             transaction.MCC,
			Currency:        transaction.Currency,
			Amount:          transaction.Amount,
			SGDAmount:       transaction.SGDAmount,
			TransactionID:   transaction.TransactionID,
			TransactionDate: transaction.TransactionDate,
			CardPAN:         transaction.CardPAN,
			CardType:        transaction.CardType,
			RewardAmount:    delta,
			Remarks:         remarks,
		}
		err := models.RewardCreate(reward)

		if err != nil {
			return err
		}
	} else {
		campaign := getMatchingCampaign(transaction.CardType)
		// Should not be since we will have base campaign, can consider throwing error(?)
		if campaign == nil {
			return nil // Reconsider what to return here
		}

		delta := calculateDeltaType(campaign.RewardAmount, transaction.Amount)
		remarks := "" // TODO: Map campaigns to appropriate remarks

		reward := models.Reward{
			ID:              gocql.UUID(transaction.ID),
			CardID:          gocql.UUID(transaction.CardID),
			Merchant:        transaction.Merchant,
			MCC:             transaction.MCC,
			Currency:        transaction.Currency,
			Amount:          transaction.Amount,
			SGDAmount:       transaction.SGDAmount,
			TransactionID:   transaction.TransactionID,
			TransactionDate: transaction.TransactionDate,
			CardPAN:         transaction.CardPAN,
			CardType:        transaction.CardType,
			RewardAmount:    delta,
			Remarks:         remarks,
		}
		err := models.RewardCreate(reward)

		if err != nil {
			return err
		}
	}
	return nil
}

func isExcluded(transactionDate time.Time, mcc int) bool {
	today := time.Now()
	for _, ex := range ExclusionsEtcd {
		if ex.MCC == mcc && ex.ValidFrom.Before(today) {
			return true
		}
	}
	return false
}

func getMatchingCampaign(cardType string) (campaign *models.Campaign) {
	// TODO: @oversparkling - change current time to transaction date
	today := time.Now()
	for _, campaign := range CampaignsEtcd {
		if campaign.RewardProgram == cardType && campaign.Start.Before(today) && campaign.End.After(today) {
			return &campaign
		}
	}
	return nil
}

func calculateDeltaType(rewardAmount int, spentAmount float64) float64 {
	// TODO: include other logic here
	// Min spend
	return float64(rewardAmount) * spentAmount
}
