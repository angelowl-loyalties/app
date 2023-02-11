package main

import (
	"fmt"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/internal"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/models"
)

func main() {
	var transaction = models.Seed_transaction
	fmt.Println(transaction)
	internal.ProcessMessage(transaction)

}
