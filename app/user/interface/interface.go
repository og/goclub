package IUserService
import (
	m "github.com/og/goclub/app/model"
	vd "github.com/og/juice/validator"
)

type Interface interface {
	SignIn (req ReqSignIn)(reply ReplySignIn, reject error)
}
type ReqSignIn struct {
	UserName string `json:"userName"`
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	SmsAuthCode string `json:"smsAuthCode"`
}
func (v ReqSignIn) VD(r *vd.Rule) {
	r.String(v.UserName, vd.StringSpec{Name:"用户名"})
	r.String(v.Mobile, vd.StringSpec{Name:"手机号", Ext: []vd.StringSpec{vd.ChinaMobile()}})
	r.String(v.Password, vd.StringSpec{Name:"密码", MinRuneLen: 8, MaxRuneLen: 40,})
	r.String(v.SmsAuthCode, vd.StringSpec{Name:"短信验证码",MinRuneLen: 4, MaxRuneLen:4})
}
type ReplySignIn struct {
	UserID m.IDUser `json:"userID"`
}
