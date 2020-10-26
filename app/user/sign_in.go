package userSerive

import (
	m "github.com/og/goclub/app/model"
	ISmsService "github.com/og/goclub/app/sms/interface"
	IUserDataStorage "github.com/og/goclub/app/user/data_storage/interface"
	IUserService "github.com/og/goclub/app/user/interface"
)
func (dep Service) SignIn(req IUserService.ReqSignIn)(reply IUserService.ReplySignIn, reject error) {
	reject = dep.vdU.Check(req) ; if reject != nil { return }
	var hasUser bool
	_, hasUser = dep.userDS.UserByName(req.UserName)
	if hasUser {
		reject = dep.resU.Reject("用户名已存在", false)
		return
	}
	_, hasUser = dep.userDS.UserByMobile(req.Mobile)
	if hasUser {
		reject = dep.resU.Reject("手机号已存在", false)
		return
	}
	reject = dep.smsS.VerifyAuthCode(ISmsService.ReqVerifyAuthCode{
		Kind: ISmsService.AuthCodeKind(nil).Enum().SignIn,
		Mobile: req.Mobile,
		AuthCode:req.Mobile,
	}) ; if reject != nil {return }
	var user m.User
	user, reject = dep.userDS.CreateUser(IUserDataStorage.ReqCreateUser{
		Name: req.UserName,
		Mobile: req.Mobile,
		Password: req.Password,
	}) ; if reject != nil {return}
	reply.UserID = user.ID
	return
}

