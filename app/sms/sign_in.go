package smsService

import (
	ISmsDataStorage "github.com/og/goclub/app/sms/data_storage/interface"
	ISmsService "github.com/og/goclub/app/sms/interface"
)

const verifySignInAuthCodeFailLimit = 10
func (dep Service) VerifySignInAuthCode(req ISmsService.ReqVerifyAuthCode) (reject error) {
	authData := ISmsDataStorage.AuthCodeData{
		Kind: ISmsDataStorage.AuthCodeKind(req.Kind.String()),
		Mobile: req.Mobile,
	}
	authCode, hasAuthCode, reject := dep.smsDS.AuthCode(authData) ; if reject != nil {return}
	if hasAuthCode {
		if req.AuthCode == authCode {
			// 成功后需让已发送的短信验证码失效
			reject = dep.smsDS.DeleteAuthCode(authData) ; if reject != nil {return}
			return nil
		} else {
			// 当错误到一定的次数时，让已发送的短信验证码失效
			reject = dep.smsDS.IncrAuthCodeVerifyFailCount(authData) ; if reject != nil {return}
			failCount, reject := dep.smsDS.AuthCodeVerifyFailCount(authData); if reject != nil {return}
			if failCount >= verifySignInAuthCodeFailLimit {
				reject = dep.smsDS.DeleteAuthCode(authData) ; if reject != nil {return}
				reject = dep.smsDS.ResetAuthCodeVerifyFailCount(authData) ; if reject != nil {return}
				return dep.resU.Reject("短信验证码错误次数过多，请重新发送", false)
			}
			return dep.resU.Reject("短信验证码错误", false)
		}
	} else {
		return dep.resU.Reject("短信验证码失效或不存在请重新发送", false)
	}
}