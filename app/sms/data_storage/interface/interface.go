package ISmsDataStorage

import vd "github.com/og/juice/validator"

type Interface interface {
	AuthCode(authData AuthCodeData)(authCode string, hasAuthCode bool, reject error)
	DeleteAuthCode(authData AuthCodeData) (reject error)
	AuthCodeVerifyFailCount(authData AuthCodeData) (failCount uint,reject error)
	IncrAuthCodeVerifyFailCount(authData AuthCodeData) (reject error)
	ResetAuthCodeVerifyFailCount(authData AuthCodeData) (reject error)
}
type AuthCodeData struct {
	Kind AuthCodeKind
	Mobile string
}
func (v AuthCodeData) VD(r *vd.Rule) {
	r.String(v.Mobile, vd.StringSpec{Name:"手机号", Ext: []vd.StringSpec{vd.ChinaMobile()}})
	r.String(v.Kind.String(), vd.StringSpec{Name:"类型", Enum: vd.EnumValues(v.Kind.Enum())})
}
type AuthCodeKind string
func NewAuthCodeKind (s string) AuthCodeKind {
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