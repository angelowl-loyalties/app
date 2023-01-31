package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/db"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"net/http"
)

// GetTransactions - GET /transaction
// Get all transactions
func GetAllTransactions(c *gin.Context) {
	var transaction []models.Transaction
	m := map[string]interface{}{}

	iter := db.DB.Query("SELECT * FROM transaction").Iter()
	for iter.MapScan(m) {
		transaction = append(transaction, models.Transaction{
			ID: m["id"].(gocql.UUID),
			CardID: m["card_id"].(gocql.UUID),
			Merchant: m["merchant"].(string),
			MCC: m["mcc"].(int),
			Currency: m["currency"].(string),
			Amount: m["amount"].(float64),
			SGDAmount: m["sgd_amount"].(float64),
			TransactionID: m["transaction_id"].(string),  
			TransactionDate: m["transaction_date"].(string),
			CardPAN: m["card_pan"].(string),
			CardType: m["card_type"].(string),
		})
		m = map[string]interface{}{}
	}

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}
