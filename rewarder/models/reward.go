package models

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
)

type Reward struct {
	ID              gocql.UUID `json:"id"`               // cassandra uuid
	CardID          gocql.UUID `json:"card_id"`          // cassandra uuid
	Merchant        string     `json:"merchant"`         // cassandra text
	MCC             int        `json:"mcc"`              // cassandra int
	Currency        string     `json:"currency"`         // cassandra text
	Amount          float64    `json:"amount"`           // cassandra double
	SGDAmount       float64    `json:"sgd_amount"`       // cassandra double
	TransactionID   string     `json:"transaction_id"`   // cassandra text
	TransactionDate string     `json:"transaction_date"` // cassandra text
	CreatedAt       string     `json:"created_at"`       // cassandra text
	CardPAN         string     `json:"card_pan"`         // cassandra text
	CardType        string     `json:"card_type"`        // cassandra text
	RewardAmount    float64    `json:"reward_amount"`    // cassandra double
	Remarks         string     `json:"remarks"`          // cassandra text
}

func RewardCreate(reward Reward) error {
	stmt, names := qb.Insert("angelowl.rewards").Columns(
		"id",
		"card_id",
		"merchant",
		"mcc",
		"currency",
		"amount",
		"sgd_amount",
		"transaction_id",
		"transaction_date",
		"created_at",
		"card_pan",
		"card_type",
		"reward_amount",
		"remarks").ToCql()

	q := gocqlx.Query(DB.Query(stmt), names).BindStruct(&reward)

	err := q.ExecRelease()
	if err != nil {
		return err
	}

	return nil
}

// Batched Insert code. Not working

// const (
// 	MaxBatchSize     = 2000
// 	MaxBatchWaitTime = 100 * time.Millisecond
// )

// var (
// 	rewardBatch []Reward
// 	rewardChan  = make(chan Reward, MaxBatchSize)
// 	rewardMutex = &sync.Mutex{}
// 	batchTimer  *time.Timer
// )

// func init() {
// 	batchTimer = time.NewTimer(MaxBatchWaitTime)
// 	go batchInsertLoop()
// }

// func batchInsertLoop() {
// 	for {
// 		select {
// 		case reward := <-rewardChan:
// 			addRewardToBatch(reward)
// 		case <-batchTimer.C:
// 			flushRewardBatch()
// 			resetBatchTimer()
// 		}
// 	}
// }

// func resetBatchTimer() {
// 	if !batchTimer.Stop() {
// 		<-batchTimer.C
// 	}
// 	batchTimer.Reset(MaxBatchWaitTime)
// }

// func addRewardToBatch(reward Reward) {
// 	rewardMutex.Lock()
// 	defer rewardMutex.Unlock()
// 	rewardBatch = append(rewardBatch, reward)

// 	if len(rewardBatch) >= MaxBatchSize {
// 		flushRewardBatch()
// 		resetBatchTimer()
// 	}
// }

// func flushRewardBatch() {
// 	rewardMutex.Lock()
// 	defer rewardMutex.Unlock()

// 	if len(rewardBatch) == 0 {
// 		return
// 	}

// 	batch := DB.NewBatch(gocql.LoggedBatch)

// 	stmt, _ := qb.Insert("angelowl.rewards").Columns(
// 		"card_id",
// 		"transaction_date",
// 		"id",
// 		"card_pan",
// 		"card_type",
// 		"amount",
// 		"created_at",
// 		"currency",
// 		"mcc",
// 		"merchant",
// 		"remarks",
// 		"reward_amount",
// 		"sgd_amount",
// 		"transaction_id",
// 	).ToCql()

// 	for _, reward := range rewardBatch {
// 		batch.Query(stmt, reward.CardID, reward.TransactionDate, reward.ID, reward.CardPAN, reward.CardType,
// 			reward.Amount, reward.CreatedAt, reward.Currency, reward.MCC, reward.Merchant, reward.Remarks,
// 			reward.RewardAmount, reward.SGDAmount, reward.TransactionID,
// 		)
// 	}

// 	err := DB.ExecuteBatch(batch)
// 	if err != nil {
// 		log.Printf("Error executing batch insert: %v", err)
// 	}

// 	rewardBatch = rewardBatch[:0]
// }

// func RewardCreate(reward Reward) error {
// 	rewardChan <- reward
// 	return nil
// }
