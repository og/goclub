package topicService

import (
	m "github.com/og/goclub/app/model"
	ITopicDataStorage "github.com/og/goclub/app/topic/data_storage/interface"
	ITopicService "github.com/og/goclub/app/topic/interface"
	respU "github.com/og/goclub/app/util/response"
	vdU "github.com/og/goclub/app/util/validator"
)

func (dep Service) Create(data ITopicService.ReqCreateTopic) (topicID m.IDTopic, reject error) {
	reject = vdU.Check(data) ; if reject != nil { return }
	_, hasTopic, reject := dep.topicDS.TopicByTitle(data.Title) ; if reject != nil {return}
	if hasTopic {
		reject = respU.Reject(data.Title + "已经存在", false) ; return
	}
	topic, reject := dep.topicDS.CreateTopic(ITopicDataStorage.CreateTopic{
		Title: data.Title,
		CoverPhoto: data.CoverPhoto,
		Intro: data.Intro,
	}) ; if reject != nil {return}
	return topic.ID, nil
}
