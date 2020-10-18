package ISmsService

import (
	ISmsDataStorage "github.com/og/goclub/app/sms/data_storage/interface"
	vd "github.com/og/juice/validator"
)

type Interface interface {
	SendAuthCode(req ReqSendAuthCode) (reject error)
	VerifyAuthCode(req ReqVerifyAuthCode) (reject error)
}

type ReqSendAuthCode struct {
	Kind ReqVerifyAuthCodeKind
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
	Kind ReqVerifyAuthCodeKind
	Mobile string
	AuthCode string
}
type ReqVerifyAuthCodeKind string

func (v ReqVerifyAuthCodeKind) String() string {
	return string(v)
}
func (ReqVerifyAuthCodeKind) Enum() (e struct{
	Login ReqVerifyAuthCodeKind
	SignIn ReqVerifyAuthCodeKind
}) {
	e.Login = "login"
	e.SignIn = "signIn"
	return
}
func (kind ReqVerifyAuthCodeKind) AuthDataKind () ISmsDataStorage.AuthCodeKind {
	return ISmsDataStorage.NewAuthCodeKind(kind.String())
}