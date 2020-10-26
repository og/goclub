package ISmsService

import (
	"errors"
	ISmsDataStorage "github.com/og/goclub/app/sms/data_storage/interface"
	vd "github.com/og/juice/validator"
)

type Interface interface {
	SendAuthCode(req ReqSendAuthCode) (reject error)
	VerifyAuthCode(req ReqVerifyAuthCode) (reject error)
}

type ReqSendAuthCode struct {
	Kind AuthCodeKind
	Mobile string
}

func (v ReqSendAuthCode) VD(r *vd.Rule) {
	r.String(v.Mobile, vd.StringSpec{Name:"手机号", Ext: []vd.StringSpec{vd.ChinaMobile()}})
	r.String(v.Kind.String(), vd.StringSpec{
		Name:"类型",
		Enum: vd.EnumValues(v.Kind.Enum()),
	})
}
type ReqVerifyAuthCode struct {
	Kind AuthCodeKind
	Mobile string
	AuthCode string
}
type AuthCodeKind string
func NewAuthCodeKind(s string) AuthCodeKind {
	return AuthCodeKind(s)
}
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
		panic(errors.New("AuthCodeKind kind error"))
	case enum.Login:
		Login(0)
	case enum.SignIn:
		SignIn(false)
	}
}
func (v AuthCodeKind) SmsDataStorageAuthCodeKind() (kind ISmsDataStorage.AuthCodeKind) {
	v.Switch(func(_login int) {
		kind = kind.Enum().Login
	}, func(_signIn bool) {
		kind = kind.Enum().SignIn
	})
	return
}