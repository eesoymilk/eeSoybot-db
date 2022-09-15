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
	ID       int64 `bun:"id,pk"`
	GuildID  int64
	Name     string `bun:",notnull"`
	Animated bool   `bun:",notnull"`
	Custom   bool   `bun:",notnull"`
}

type TriggerBase struct {
	ID      int64   `bun:"id,pk"`
	GuildID int64   `bun:",notnull"`
	chance  float64 `bun:",notnull"`
}

type AutoReplyTrigger struct {
	TriggerBase `bun:",extend"`
	AutoReplyID int64 `bun:",notnull"`
}

type AutoReplyUserTrigger struct {
	AutoReplyTrigger `bun:",extend"`
	UserID           int64 `bun:",notnull"`
}

type AutoReplyKeywordTrigger struct {
	AutoReplyTrigger `bun:",extend"`
	Keywords         []string `bun:",array,notnull"`
}

type AutoReply struct {
	ID   int64    `bun:"id,pk"`
	Pool []string `bun:",array,notnull"`
}

type AutoReactTrigger struct {
	TriggerBase `bun:",extend"`
	AutoReactID int64 `bun:",notnull"`
}

type AutoReactUserTrigger struct {
	AutoReactTrigger `bun:",extend"`
	UserID           int64 `bun:",notnull"`
}

type AutoReactKeywordTrigger struct {
	AutoReactTrigger `bun:",extend"`
	Keyword          []string `bun:",array"`
}

type AutoReact struct {
	ID      int64 `bun:"id,pk"`
	EmojiID int64 `bun:",notnull"`
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
