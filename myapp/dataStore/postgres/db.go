package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgres_host     = "dpg-chos5indvk4goeshfcmg-a.singapore-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "postgres_admin"
	postgres_password = "wAiGDMcudXssqe80lUrLhQSIrFWALCNV"
	postgres_dbname   = "my_db_cowf"
)

var Db *sql.DB

// init() is always called before main()
func init() {
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)

	var err error

	Db, err = sql.Open("postgres", db_info)

	if err != nil {
		panic(err)
	} else {
		log.Println("Database successfully configured")
	}
}
