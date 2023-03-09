package models

import (
	"log"
	"strconv"

	"github.com/gocql/gocql"
)

var DB *gocql.Session

func InitDB(dbHost, dbPort, keyspace, table, username, password string, useSSL, createIndex bool) {
	cluster := gocql.NewCluster(dbHost)

	dbPortInt, err := strconv.Atoi(dbPort)
	if err == nil {
		cluster.Port = dbPortInt
	}

	if useSSL {
		cluster.SslOpts = &gocql.SslOptions{
			CaPath: "/root-ca.crt",
		}
	}

	if username != "" && password != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: username,
			Password: password,
		}
	}

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

	if createIndex {
		err = session.Query("CREATE INDEX IF NOT EXISTS rewards_idx ON " +
			keyspace + "." + table + " ( card_id );").Exec()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Created index on: card_id")
	}
}

func ConnectDB(dbHost, dbPort, username, password, keyspace string, useSSL bool) {
	cluster := gocql.NewCluster(dbHost)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.LocalQuorum

	dbPortInt, err := strconv.Atoi(dbPort)
	if err == nil {
		cluster.Port = dbPortInt
	}

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}

	if useSSL {
		cluster.SslOpts = &gocql.SslOptions{
			CaPath: "/root-ca.crt",
		}
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to Rewards DB")
	DB = session
}

