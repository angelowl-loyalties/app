package internal

import (
	"encoding/json"
	"log"
	"strconv"
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

	transactionDate, err := time.Parse(YYYYMMDD, transaction.TransactionDate)
	if err != nil {
		log.Println(err)
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
			log.Println(err)
			return err
		}

	} else {
		campaigns := getMatchingCampaigns(transaction)
		baseDelta := 0.0
		promoDelta := 0.0
		remarks := ""
		baseMatchedCampaigns := campaigns[0]
		promoMatchedCampaigns := campaigns[1]

		if transaction.Currency != "SGD" {
			transaction.Amount = convertToSGD(transaction.Amount)
		}
		for _, campaign := range baseMatchedCampaigns {

			tempDelta := calculateDeltaType(campaign.RewardAmount, transaction.Amount)
			if tempDelta > baseDelta {
				baseDelta = tempDelta
				remarks = campaign.Name + " applied"
			}
		}

		for _, campaign := range promoMatchedCampaigns {

			tempDelta := calculateDeltaType(campaign.RewardAmount, transaction.Amount)
			if tempDelta > promoDelta {
				promoDelta = tempDelta
				remarks = campaign.Name + " applied"
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
			RewardAmount:    baseDelta + promoDelta,
			Remarks:         remarks,
		}
		err := models.RewardCreate(reward)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func isExcluded(transactionDate time.Time, mcc int) bool {
	if mcc < 1 || mcc > 9999 {
		return true
	}

	for _, ex := range ExclusionsEtcd {
		if ex.MCC == mcc && ex.ValidFrom.Before(transactionDate) {
			return true
		}
	}
	return false
}

func getMatchingCampaigns(transaction models.Transaction) (campaign [][]models.Campaign) {
	//Returns a 2D array of campaigns, [ [BaseMatchedCampaigns], [PromoMatchedCampaigns] ]

	promoMatchingCampaigns := []models.Campaign{}
	baseMatchingCampaigns := []models.Campaign{}

	resultMatchingCampaigns := [][]models.Campaign{}
	for _, campaign := range BaseCampaignsEtcd {

		if isCampaignMatch(campaign, transaction) {
			baseMatchingCampaigns = append(baseMatchingCampaigns, campaign)
		}
	}

	for _, campaign := range PromoCampaignsEtcd {
		if isCampaignMatch(campaign, transaction) {
			promoMatchingCampaigns = append(promoMatchingCampaigns, campaign)
		}
	}

	resultMatchingCampaigns = append(resultMatchingCampaigns, baseMatchingCampaigns)
	resultMatchingCampaigns = append(resultMatchingCampaigns, promoMatchingCampaigns)
	return resultMatchingCampaigns
}

func isCampaignMatch(campaign models.Campaign, transaction models.Transaction) bool {
	transactionDate, err := time.Parse(YYYYMMDD, transaction.TransactionDate)

	//Should not reach here since its second time parsing
	if err != nil {
		log.Println(err)
		return false
	}
	if campaign.RewardProgram != transaction.CardType {
		return false
	}
	if !campaign.Start.Before(transactionDate) || !campaign.End.After(transactionDate) {
		return false
	}
	if campaign.ForForeignCurrency && transaction.Currency == "SGD" {
		return false
	}
	if campaign.Merchant != "" && !strings.Contains(transaction.Merchant, campaign.Merchant) {
		return false
	}
	if campaign.MinSpend > transaction.Amount {
		return false
	}
	campaignMcc := strings.Split(campaign.MCC, ",")
	for _, mcc := range campaignMcc {
		intMCC, err := strconv.Atoi(mcc)
		if err != nil {
			// log.Fatalln(err)
			return false
		}
		if intMCC == 0 || intMCC == transaction.MCC {
			return true
		}
	}
	return false
}

func calculateDeltaType(rewardAmount int, spentAmount float64) float64 {

	return float64(rewardAmount) * spentAmount
}

func convertToSGD(spendAmount float64) float64 {
	return spendAmount * 1.34
}
