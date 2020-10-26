package smsService

import (
	ISmsDataStorage "github.com/og/goclub/app/sms/data_storage/interface"
	ISmsService "github.com/og/goclub/app/sms/interface"
	grand "github.com/og/x/rand"
	"log"
)

func (dep Service) SendAuthCode(req ISmsService.ReqSendAuthCode) (reject error) {
	reject = dep.vdU.Check(req) ; if reject != nil {return}
	authCode := string(grand.RunesBySeed("0123456789", 4))
	log.Print("Send sms: auth code is " + authCode)
	reject = dep.smsDS.CreateAuthCodeRecord(ISmsDataStorage.CreateAuthCodeRecord{
		Kind:     req.Kind.SmsDataStorageAuthCodeKind(),
		Mobile:   req.Mobile,
		AuthCode: authCode,
	}) ; if reject != nil {return}
	return nil
}
