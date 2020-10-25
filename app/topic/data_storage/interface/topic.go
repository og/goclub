package ITopicDataStorage

import m "github.com/og/goclub/app/model"

type Interface interface {
	TopicByTitle(title string) (topic m.Topic, hasTopic bool, reject error)
	CreateTopic(data CreateTopic) (topic m.Topic, reject error)
	TopicByID(id m.IDTopic) (topic m.Topic, hasTopic bool, reject error)
	// MustHasTopicByID 封装了 TopicByID ，如果 hasTopic == false  时 reject != nil
	MustHasTopicByID(id m.IDTopic) (topic m.Topic, reject error)
	UpdateTopicIntro(topicID m.IDTopic, intro string) (reject error)
	UpdateTopicCoverPhoto(topicID m.IDTopic, coverPhoto string) (reject error)
	CreateTopicLogCoverPhoto(data CreateTopicLogCoverPhoto) (reject error)
	CreateTopicLogIntro(data CreateTopicLogIntro) (reject error)
}

type CreateTopic struct {
	Title string
	CoverPhoto string
	Intro string
}
type CreateTopicLogCoverPhoto struct {
	OldCoverPhoto string
	NewCoverPhoto string
}
type CreateTopicLogIntro struct {
	OldIntro string
	NewIntro string
}