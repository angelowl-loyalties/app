package internal

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/models"
	"github.com/gocql/gocql"
)

const (
	YYYYMMDD = "2006-01-02"
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
	transactionDate, err := time.Parse(YYYYMMDD, transaction.TransactionDate)

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
			log.Fatalln(err)
			return err
		}

	} else {
		campaigns := getMatchingCampaigns(transaction)

		if campaigns != nil {
			delta := 0.0
			remarks := ""
			for _, campaign := range campaigns {
				tempDelta := calculateDeltaType(campaign.RewardAmount, transaction.Amount)
				if tempDelta > delta {
					delta = tempDelta
					remarks = campaign.Name + " applied" //TODO: Format this string to appropriate description
				}
			}

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
				log.Fatalln(err)
				return err
			}
		} else {
			//Should never reach here
			log.Fatalln("No matching campaigns found")
		}

		// } else {
		// 	delta := 69.0
		// 	remarks := "Base campaign applied"

		// 	reward := models.Reward{
		// 		ID:              gocql.UUID(transaction.ID),
		// 		CardID:          gocql.UUID(transaction.CardID),
		// 		Merchant:        transaction.Merchant,
		// 		MCC:             transaction.MCC,
		// 		Currency:        transaction.Currency,
		// 		Amount:          transaction.Amount,
		// 		SGDAmount:       transaction.SGDAmount,
		// 		TransactionID:   transaction.TransactionID,
		// 		TransactionDate: transaction.TransactionDate,
		// 		CardPAN:         transaction.CardPAN,
		// 		CardType:        transaction.CardType,
		// 		RewardAmount:    delta,
		// 		Remarks:         remarks,
		// 	}
		// 	err := models.RewardCreate(reward)
		// 	// TODO: Proper Error handling
		// 	if err != nil {
		// 		log.Fatalln(err)
		// 		return err
		// 	}
		// }
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

func getMatchingCampaigns(transaction models.Transaction) (campaign []*models.Campaign) {
	// TODO: start using the transactions start date instead of time now
	// transactionDate, _ := time.Parse(YYYYMMDD, transactionDateStr)
	var matchingCampaigns []*models.Campaign
	for _, campaign := range CampaignsEtcd {
		if isCampaignMatch(&campaign, &transaction) {
			matchingCampaigns = append(matchingCampaigns, &campaign)
		}
	}
	return matchingCampaigns
}

func isCampaignMatch(campaign *models.Campaign, transaction *models.Transaction) bool {
	if campaign.RewardProgram != transaction.CardType {
		return false
	}

	//TODO: Change to proper transaction date
	if !campaign.Start.Before(time.Now()) || !campaign.End.After(time.Now()) || campaign.MinSpend > transaction.Amount {
		return false
	}
	if campaign.IsForeign && transaction.Currency == "SGD" {
		return false
	}
	if campaign.ValidName != "" && strings.Contains(transaction.Merchant, campaign.ValidName) {
		return false
	}
	return true
}

func calculateDeltaType(rewardAmount int, spentAmount float64) float64 {
	// TODO: include other logic here
	// Min spend
	return float64(rewardAmount) * spentAmount
}

func convertToSGD(spendAmount float64) float64 {

	// TODO: Change to proper USD handling
	return spendAmount * 1.34
}
