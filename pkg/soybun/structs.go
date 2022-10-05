package soybun

import "github.com/uptrace/bun"

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
