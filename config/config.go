package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func LoadConfig() {
	dns := "user=postgres.bwpxemxjvkrxeroqohah password=11_h@md@1234 host=aws-0-ap-southeast-1.pooler.supabase.com port=6543 dbname=postgres"
	var err error
	DB, err = sql.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		DB.Close()
		log.Fatal(err)
	}
}
