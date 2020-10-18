package smsService

import (
	ISmsDataStorage "github.com/og/goclub/app/sms/data_storage/interface"
	responseUtil "github.com/og/goclub/app/util/response"
)

type Service struct {
	smsDS ISmsDataStorage.Interface
	resU responseUtil.Util
}
