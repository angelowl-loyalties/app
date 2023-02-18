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

	err := json.Unmarshal([]byte(messageJSON), &transaction) // Convert JSON message to transaction object
	if err != nil {
		// log.Fatalln(err)
		return nil
	}

	return ProcessMessage(transaction)
}

func ProcessMessage(transaction models.Transaction) error {
	//TODO: Proper error handling
	transactionDate, err := time.Parse(layout, transaction.TransactionDate)

	// TODO: Proper Error handling
	if err != nil {
		// log.Fatalln(err)
		return nil //TODO: Fix this to do something when date couldnt be parsed
	}

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
		// TODO: Proper Error handling
		if err != nil {
			// log.Fatalln(err)
			return err
		}

	} else {
		campaign := getMatchingCampaign(transaction.CardType, transaction.TransactionDate)

		if campaign == nil {
			return nil // TODO: Reconsider what to return here
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
		// TODO: Proper Error handling
		if err != nil {
			// log.Fatalln(err)
			return err
		}

	}
	return nil
}

func isExcluded(transactionDate time.Time, mcc int) bool {
	// TODO: Replace with transactionDate
	today := time.Now()
	for _, ex := range ExclusionsEtcd {
		if ex.MCC == mcc && ex.ValidFrom.Before(today) {
			return true
		}
	}
	return false
}

func getMatchingCampaign(cardType string, transactionDateStr string) (campaign *models.Campaign) {
	// TODO: start using the transactions start date instead of time now
	// transactionDate, _ := time.Parse(layout, transactionDateStr)
	transactionDate := time.Now()
	for _, campaign := range CampaignsEtcd {
		if campaign.RewardProgram == cardType && campaign.Start.Before(transactionDate) && campaign.End.After(transactionDate) {
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
