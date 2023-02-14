package internal

import (
	"fmt"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/models"
)

func ProcessMessage(transaction models.Transaction) {
	//TODO: Consume from etcd and MQ
	campaigns := models.Seed_campaigns
	exclusions := models.Seed_exclusions
	// var campaigns []models.Campaign
	// var exclusions []models.Exclusion
	layout := "02-01-06"
	transactionDate, _ := time.Parse(layout, transaction.TransactionDate)
	if isExcluded(transactionDate, transaction.MCC, exclusions) {
		delta := 0
		remark := "Campaigns don't apply"
		fmt.Println(delta, remark)
	} else {
		campaign := getMatchingCampaign(transaction.CardType, campaigns)
		//Should not be since we will have base campaign, can consider throwing error(?)
		if campaign == nil {
			return
		}
		fmt.Println(campaign.RewardAmount)
		delta := calculateDeltaType(campaign.RewardAmount, transaction.Amount)
		fmt.Println(delta)
	}
}

func isExcluded(transactionDate time.Time, mcc int, exclusions []models.Exclusion) bool {
	//chang eto txn date
	today := time.Now()
	for _, ex := range exclusions {
		if ex.MCC == mcc && ex.ValidFrom.Before(today) {
			return true
		}
	}
	return false
}

func getMatchingCampaign(cardType string, campaigns []models.Campaign) (campaign *models.Campaign) {
	//Change today to transaction date
	today := time.Now()
	for _, campaign := range campaigns {
		if campaign.CardType == cardType && campaign.Start.Before(today) && campaign.End.After(today) {
			return &campaign
		}
	}
	return nil
}

func calculateDeltaType(rewardAmount int, spentAmount float64) float64 {
	//To include other logic here
	// Min spend
	return float64(rewardAmount) * spentAmount
}
