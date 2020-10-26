package m

import (
	"database/sql"
	"errors"
	"time"
)

type  IDTopic string
func (id IDTopic) String() string {return string(id)}
type Topic struct {
	ID IDTopic `db:"id"`
	Title string `db:"title"`
	CoverPhoto string `db:"cover_photo"`
	Intro string `db:"intro"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type IDQuestionRelationTopic string
type QuestionRelationTopic struct {
	ID IDQuestion `db:"id"`
	QuestionID IDQuestion `db:"question_id"`
	TopicID IDTopic `db:"topic_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type IDTopicLog string
type TopicLog struct {
	ID TopicLog `db:"id"`

	TopicID IDTopic `db:"topic_id"`
	UserID IDUser `db:"user_id"`
	Kind TopicLogKind `db:"kind"`

	OldCoverPhoto string `db:"old_cover_photo"`
	NewCoverPhoto string `db:"new_cover_photo"`

	OldIntro string `db:"old_intro"`
	NewIntro string `db:"new_intro"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
type TopicLogKind string
func (kind TopicLogKind) String() string {
	return string(kind)
}
func (TopicLogKind) Enum () (e struct{
	CoverPhoto TopicLogKind
	Intro TopicLogKind
}) {
	e.CoverPhoto = "coverPhoto"
	e.Intro = "intro"
	return
}
func (v TopicLogKind) Switch(
	CoverPhoto func(_CoverPhoto int),
	Intro func(_Intro bool),
) {
	enum := v.Enum()
	switch v {
	default:
		panic(errors.New("TopicLogKind err (" + v.String() +")"))
	case enum.CoverPhoto:
		CoverPhoto(0)
	case enum.Intro:
		Intro(false)
	}
}