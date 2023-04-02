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
		log.Printf("Error unmarshalling transaction JSON, skipping: %v", err)
		return nil
	}

	return ProcessMessage(transaction)
}

func ProcessMessage(transaction models.Transaction) error {
	transactionDate, err := time.Parse(YYYYMMDD, transaction.TransactionDate)
	if err != nil {
		log.Printf("Error parsing transaction date, skipping: %v", err)
		return nil
	}

	if IsExcluded(transactionDate, transaction.MCC) {
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
			CreatedAt:       time.Now().Format(YYYYMMDD),
			CardPAN:         maskCardPAN(transaction.CardPAN),
			CardType:        transaction.CardType,
			RewardAmount:    delta,
			Remarks:         remarks,
		}

		err := models.RewardCreate(reward)
		if err != nil {
			log.Printf("Error creating reward: %v", err)
			return err
		}

	} else {
		campaigns := GetMatchingCampaigns(transaction)
		baseDelta := 0.0
		promoDelta := 0.0
		remarks := ""
		baseMatchedCampaigns := campaigns[0]
		promoMatchedCampaigns := campaigns[1]

		if transaction.Currency != "SGD" {
			transaction.Amount = ConvertToSGD(transaction.Amount)
		}
		for _, campaign := range baseMatchedCampaigns {

			tempDelta := CalculateDeltaType(campaign.RewardAmount, transaction.Amount)
			if tempDelta > baseDelta {
				baseDelta = tempDelta
				remarks = campaign.Name + " applied"
			}
		}

		for _, campaign := range promoMatchedCampaigns {

			tempDelta := CalculateDeltaType(campaign.RewardAmount, transaction.Amount)
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
			CreatedAt:       time.Now().Format(YYYYMMDD),
			CardPAN:         maskCardPAN(transaction.CardPAN),
			CardType:        transaction.CardType,
			RewardAmount:    baseDelta + promoDelta,
			Remarks:         remarks,
		}
		err := models.RewardCreate(reward)
		if err != nil {
			log.Printf("Error creating reward: %v", err)
			return err
		}
	}

	return nil
}

func IsExcluded(transactionDate time.Time, mcc int) bool {
	RefreshFromEtcd()
	if mcc < 1 || mcc > 9999 {
		return true
	}

	exclusionsMutex.RLock()
	for _, ex := range ExclusionsEtcd {
		if ex.MCC == mcc && ex.ValidFrom.Before(transactionDate) {
			return true
		}
	}
	exclusionsMutex.RUnlock()

	return false
}

func GetMatchingCampaigns(transaction models.Transaction) (campaign [][]models.Campaign) {
	RefreshFromEtcd()
	//Returns a 2D array of campaigns, [ [BaseMatchedCampaigns], [PromoMatchedCampaigns] ]

	var baseMatchingCampaigns []models.Campaign
	var promoMatchingCampaigns []models.Campaign
	var resultMatchingCampaigns [][]models.Campaign

	baseCampaignMutex.RLock()
	for _, campaign := range BaseCampaignsEtcd {
		if IsCampaignMatch(campaign, transaction) {
			baseMatchingCampaigns = append(baseMatchingCampaigns, campaign)
		}
	}
	baseCampaignMutex.RUnlock()

	promoCampaignMutex.RLock()
	for _, campaign := range PromoCampaignsEtcd {
		if IsCampaignMatch(campaign, transaction) {
			promoMatchingCampaigns = append(promoMatchingCampaigns, campaign)
		}
	}
	promoCampaignMutex.RUnlock()

	resultMatchingCampaigns = append(resultMatchingCampaigns, baseMatchingCampaigns)
	resultMatchingCampaigns = append(resultMatchingCampaigns, promoMatchingCampaigns)

	return resultMatchingCampaigns
}

func IsCampaignMatch(campaign models.Campaign, transaction models.Transaction) bool {
	transactionDate, err := time.Parse(YYYYMMDD, transaction.TransactionDate)

	if err != nil {
		log.Printf("Error parsing transaction date in campaign match, skipping: %v", err)
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
			log.Printf("Error parsing MCC: %v", err)
			return false
		}
		if intMCC == 0 || intMCC == transaction.MCC {
			return true
		}
	}

	return false
}

func CalculateDeltaType(rewardAmount int, spentAmount float64) float64 {
	return float64(rewardAmount) * spentAmount
}

func ConvertToSGD(spendAmount float64) float64 {
	return spendAmount * 1.34
}

func maskCardPAN(cardPAN string) string {
	strLen := len(cardPAN)

	mask := strings.Repeat("*", strLen-4)
	lastFour := cardPAN[strLen-4:]

	return mask + lastFour
}
