package internal

import (
	"fmt"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/models"
)

func ProcessMessage(transaction models.Transaction) {
	campaigns := models.Seed_campaigns
	exclusions := models.Seed_exclusions
	//TODO: Consume from etcd and MQ
	// var campaigns []models.Campaign
	// var exclusions []models.Exclusion
	layout := "02-01-06"
	transactionDate, _ := time.Parse(layout, transaction.TransactionDate)
	if isExcluded(transactionDate, transaction.MCC) {
		delta := 0
		remark := "Campaigns don't apply"

		// TODO: The function below adds a reward object to cassandra
		fmt.Println(delta, remark)
	} else {
		campaign := getMatchingCampaign(transaction.CardType)
		// Should not be since we will have base campaign, can consider throwing error(?)
		if campaign == nil {
			return
		}
		fmt.Println(campaign.RewardAmount)
		delta := calculateDeltaType(campaign.RewardAmount, transaction.Amount)
		fmt.Println(delta)
	}
}

func isExcluded(transactionDate time.Time, mcc int) bool {
	//TODO: View exclusions in etcd
	exclusions := models.Seed_exclusions
	//change to txn date
	today := time.Now()
	for _, ex := range exclusions {
		if ex.MCC == mcc && ex.ValidFrom.Before(today) {
			return true
		}
	}
	return false
}

func getMatchingCampaign(cardType string) (campaign *models.Campaign) {
	//TODO: View campaigns in etcd
	campaigns := models.Seed_campaigns
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
