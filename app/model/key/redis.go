package mkey

import (
	"strings"
)

type Key string
func (key Key) String() string {
	return string(key)
}
const (
	smsAuthCode Key = "sms:authCode:{{Kind}}:{{Mobile}}"
	smsAuthCodeVerifyFailCount Key = "sms:authCode:verify:failCount{{Kind}}:{{Mobile}}"
)
type QuerySmsAuthCode struct {
	Kind string
	Mobile string
}
func SmsAuthCode(query QuerySmsAuthCode) (key string, value string) {
	return strings.Join([]string{
		"sms","authCode", query.Kind, query.Mobile,
	}, ":"), ""
}
