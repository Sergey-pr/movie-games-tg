package persist

import (
	"database/sql"
	"github.com/Sergey-pr/movie-games-tg/config"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
	"log"
)

// Db is a db connection variable
var Db = connectPostgresql()

// connectPostgresql returns a db connection variable
func connectPostgresql() *goqu.Database {
	// Set db dialect
	dialect := goqu.Dialect("postgres")
	// Set db driver and connection by dsn string
	pgDb, err := sql.Open("postgres", config.AppConfig.DatabaseDns)
	if err != nil {
		log.Fatalln(err)
	}
	return dialect.DB(pgDb)
}
