package bootstrap

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"

	"notas/model"
)

func newDatabase(ctx context.Context) model.PgxPool {
	host := os.Getenv("ODL_HOST")
	port := os.Getenv("ODL_PORT")
	user := os.Getenv("ODL_USER")
	password := os.Getenv("ODL_PASS")
	dbname := os.Getenv("ODL_DB")
	sslMode := os.Getenv("ODL_SSLMODE")

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("could not parse config pgxpool, err: %v", err)
	}

	config.ConnConfig.RuntimeParams["application_name"] = model.ApplicationName
	config.ConnConfig.Tracer = otelpgx.NewTracer()

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("could not connection to db, err: %v", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("could ping database, err: %v", err)
	}

	return db
}
