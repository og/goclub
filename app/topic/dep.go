package topicService

import ITopicDataStorage "github.com/og/goclub/app/topic/data_storage/interface"

type Service struct {
	topicDS ITopicDataStorage.Interface
}