package m

import (
	"database/sql"
	"time"
)

type IDQuestion string
func (id IDQuestion) String () string { return string(id) }
type Question struct {
	ID IDQuestion `db:"id"`
	UserID IDUser `db:"user_id"`
	Title string `db:"title"`
	Content string `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
