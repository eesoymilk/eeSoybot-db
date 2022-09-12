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
)

type Guild struct {
	ID   int64  `bun:"id,pk"`
	Name string `bun:"name,notnull"`
}

type Emoji struct {
	ID       int64  `bun:"id,pk"`
	Name     string `bun:"notnull"`
	Animated bool   `bun:"notnull"`
	Custom   bool   `bun:"notnull"`
}

type AutoResponseTrigger struct {
	ID      int64 `bun:"id,pk"`
	EmojiID int64 `bun:"notnull"`
}

type AutoReact struct {
	ID      int64   `bun:"id,pk"`
	EmojiID int64   `bun:"notnull"`
	chance  float64 `bun:"notnull"`
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64  `bun:"id,pk"`
	Username      string `bun:"notnull"`
}

func main() {
	ctx := context.Background()
	router := bunrouter.New()

	router.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		fmt.Println(req.Method, req.Route(), req.Params().Map())
		return nil
	})

	// Open a PostgreSQL database.
	dsn := "postgres://postgres:510600@localhost:5432/eesoybot?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	//
	var g Guild
	err := db.NewSelect().Model((*Guild)(nil)).Where("id = ?", 874556062815100938).Scan(ctx, &g)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(g)
	}
}
