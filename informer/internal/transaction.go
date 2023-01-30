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
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	m := map[string]interface{}{}

	iter := db.DB.Query("SELECT * FROM transactions").Iter()
	for iter.MapScan(m) {
		transactions = append(transactions, models.Transaction{
			ID: m["id"].(gocql.UUID),
			// add other fields accordingly
		})
		m = map[string]interface{}{}
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}
