package topicService

import (
	ITopicDataStorage "github.com/og/goclub/app/topic/data_storage/interface"
	ITopicService "github.com/og/goclub/app/topic/interface"
	respU "github.com/og/goclub/app/util/response"
	vdU "github.com/og/goclub/app/util/validator"
)

func (dep Service) UpdateTopicIntro(data ITopicService.ReqUpdateTopicIntro) (reject error) {
	reject = vdU.Check(data) ; if reject != nil { return }
	topic, reject := dep.topicDS.MustHasTopicByID(data.TopicID) ; if reject != nil {return}
	if data.Intro != topic.Intro {
		reject = dep.topicDS.UpdateTopicIntro(data.TopicID, data.Intro) ; if reject != nil {return}
		reject = dep.topicDS.CreateTopicLogIntro(ITopicDataStorage.CreateTopicLogIntro{
			OldIntro: topic.Intro,
			NewIntro: data.Intro,
		}) ; if reject != nil {return}
	} else {
		return respU.Reject("简介没有任何修改", false)
	}
	return nil
}
func (dep Service) UpdateTopicCoverPhoto(data ITopicService.ReqUpdateTopicCoverPhoto) (reject error) {
	reject = vdU.Check(data) ; if reject != nil { return }
	topic, reject := dep.topicDS.MustHasTopicByID(data.TopicID) ; if reject != nil {return}
	if data.CoverPhoto != topic.CoverPhoto {
		reject = dep.topicDS.UpdateTopicCoverPhoto(data.TopicID, data.CoverPhoto) ; if reject != nil {return}
		reject = dep.topicDS.CreateTopicLogCoverPhoto(ITopicDataStorage.CreateTopicLogCoverPhoto{
			OldCoverPhoto: topic.CoverPhoto,
			NewCoverPhoto: data.CoverPhoto,
		}) ; if reject != nil {return}
	} else {
		return respU.Reject("封面图没有任何修改", false)
	}
	return nil
}