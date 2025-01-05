package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:",pk,autoincrement"`
	Name      string    `bun:",notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:CURRENT_TIMESTAMP"`

	Posts []*Post `bun:"rel:has-many,join:id=user_id"`
}

type Post struct {
	bun.BaseModel `bun:"table:posts"`

	ID        int64  `bun:",pk,autoincrement"`
	Title     string `bun:",notnull"`
	Content   string
	UserID    int64
	CreatedAt time.Time `bun:",nullzero,notnull,default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:CURRENT_TIMESTAMP"`

	Comments []*Comment `bun:"rel:has-many,join:id=post_id"`
}

type Comment struct {
	bun.BaseModel `bun:"table:comments"`

	ID        int64 `bun:",pk,autoincrement"`
	Content   string
	UserID    int64
	PostID    int64
	CreatedAt time.Time `bun:",nullzero,notnull,default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:CURRENT_TIMESTAMP"`
}
