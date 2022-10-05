package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bunrouter"

	"github.com/eesoymilk/eeSoybot-db/pkg/soybun"
)

func main() {
	ctx := context.Background()
	router := bunrouter.New()

	router.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		fmt.Println(req.Method, req.Route(), req.Params().Map())
		return nil
	})

	// Open a PostgreSQL database.
	dsn := "postgres://postgres:510600@192.168.51.19:5432/eesoybot?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	//
	var g soybun.Guild
	err := db.NewSelect().Model((*soybun.Guild)(nil)).Where("id = ?", 874556062815100938).Scan(ctx, &g)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(g)
	}
}
