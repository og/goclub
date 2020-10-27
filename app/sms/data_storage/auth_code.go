package smsDataStorage

import (
	"errors"
	"github.com/mediocregopher/radix/v3"
	mkey "github.com/og/goclub/app/model/key"
	ISmsDataStorage "github.com/og/goclub/app/sms/data_storage/interface"
	gconv "github.com/og/x/conv"
)


func (dep DataStorage) CreateAuthCodeRecord(record ISmsDataStorage.CreateAuthCodeRecord, expireSecond int) (reject error) {
	key,_ := mkey.SmsAuthCode(mkey.QuerySmsAuthCode{
		Kind:   record.Kind.String(),
		Mobile: record.Mobile,
	})
	reject = dep.db.Do(radix.Cmd(nil, "SET", key, record.AuthCode)) ; if reject != nil {return}
	reject = dep.db.Do(radix.Cmd(nil, "EXPIRE", key, gconv.IntString(expireSecond))) ; if reject != nil {return}
	return
}

func (dep DataStorage) AuthCode(authData ISmsDataStorage.AuthCodeData)(authCode string, hasAuthCode bool, reject error) {
	key, value := mkey.SmsAuthCode(mkey.QuerySmsAuthCode{
		Kind:   authData.Kind.String(),
		Mobile: authData.Mobile,
	})
	data := radix.MaybeNil{Rcv: &value}
	err := dep.db.Do(radix.Cmd(data, "GET", key,))
	switch {
	case err != nil:
		return "", false, err
	case data.Nil:
		return "", false, nil
	case data.Nil == false:
		return value, true, nil
	default:
		panic(errors.New("Redis GET " + key + " error"))
	}
}