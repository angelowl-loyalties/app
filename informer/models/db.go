package models

import (
	"log"

	"github.com/gocql/gocql"
)

var DB *gocql.Session

func ConnectDB(dbConnString string, keyspace string, user string, pass string) {
	cluster := gocql.NewCluster(dbConnString)
	cluster.Keyspace = keyspace
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: user,
		Password: pass,
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}

	defer session.Close()
	log.Println("Migrated Exclusion Table")

	DB = session
}
