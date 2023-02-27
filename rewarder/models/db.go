package models

import (
	"log"

	"github.com/gocql/gocql"
)

var DB *gocql.Session

func InitDB(dbConnString string, keyspace string, table string) {
	cluster := gocql.NewCluster(dbConnString)
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()

	err = session.Query("CREATE KEYSPACE IF NOT EXISTS " + keyspace +
		" WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};").Exec()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Created keyspace: ", keyspace)

	err = session.Query("CREATE TABLE IF NOT EXISTS " + keyspace + "." + table +
		" (id uuid, card_id uuid, merchant text, mcc int, currency text, amount double, sgd_amount double, " +
		"transaction_id text, transaction_date date, card_pan text, card_type text, reward_amount double, " +
		"remarks text, PRIMARY KEY(id));").Exec()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Created table: ", table)

	err = session.Query("CREATE INDEX IF NOT EXISTS rewards_idx ON " +
		keyspace + "." + table + " ( card_id );").Exec()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Created index on: card_id")
}

func ConnectDB(dbConnString string, keyspace string) {
	cluster := gocql.NewCluster(dbConnString)
	cluster.Keyspace = keyspace
	// cluster.Authenticator = gocql.PasswordAuthenticator{
	// 	Username: user,
	// 	Password: pass,
	// }

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connected to Rewards DB")

	DB = session
}