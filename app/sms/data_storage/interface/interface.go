package ISmsDataStorage

import (
	"errors"
)

type Interface interface {
	CreateAuthCodeRecord(record CreateAuthCodeRecord) (reject error)
	AuthCode(authData AuthCodeData)(authCode string, hasAuthCode bool, reject error)
	DeleteAuthCode(authData AuthCodeData) (reject error)
	AuthCodeVerifyFailCount(authData AuthCodeData) (failCount uint,reject error)
	IncrAuthCodeVerifyFailCount(authData AuthCodeData) (reject error)
	ResetAuthCodeVerifyFailCount(authData AuthCodeData) (reject error)
}
type CreateAuthCodeRecord struct {
	Kind AuthCodeKind
	Mobile string
	AuthCode string
}
type AuthCodeData struct {
	Kind AuthCodeKind
	Mobile string
}
type AuthCodeKind string
func (v AuthCodeKind) String() string {
	return string(v)
}
func (AuthCodeKind) Enum() (e struct{
	Login AuthCodeKind
	SignIn AuthCodeKind
}) {
	e.Login = "login"
	e.SignIn = "signIn"
	return
}
func (v AuthCodeKind) Switch(
	Login func(_login int),
	SignIn func(_signIn bool),
	) {
	enum := v.Enum()
	switch v {
	default:
		panic(errors.New("AuthCodeKind value error"))
	case enum.Login:
		Login(0)
	case enum.SignIn:
		SignIn(false)
	}
}