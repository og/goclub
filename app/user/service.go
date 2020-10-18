package userSerive

import (
	interfaceSmsService "github.com/og/goclub/app/sms/interface"
	IUserDataStorage "github.com/og/goclub/app/user/data_storage/interface"
	responseUtil "github.com/og/goclub/app/util/response"
	validatorUtil "github.com/og/goclub/app/util/validator"
)

type Service struct {
	userDS IUserDataStorage.Interface
	smsS   interfaceSmsService.Interface
	resU   responseUtil.Util
	vdU    validatorUtil.Util
}
