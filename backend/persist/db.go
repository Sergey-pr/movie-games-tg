package persist

import (
	"database/sql"
	"github.com/Sergey-pr/movie-games-tg/config"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"log"
)

var Db = connectPostgresql()

func connectPostgresql(dns ...string) *goqu.Database {
	if len(dns) == 0 {
		dns = append(dns, config.AppConfig.Dns)
	}

	dialect := goqu.Dialect("postgres")

	pgDb, err := sql.Open("pqHooked", dns[0])
	if err != nil {
		log.Fatalln(err)
	}

	return dialect.DB(pgDb)
}
