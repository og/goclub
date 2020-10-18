package validatorUtil

import (
	responseUtil "github.com/og/goclub/app/util/response"
	vd "github.com/og/juice/validator"
)

type Util struct {
	resU responseUtil.Util
}
var checker vd.Checker
func init () {
	checker = vd.NewCN()
}
func (dep Util) Check(data vd.Data) error {
	report := checker.Check(data)
	if report.Fail {
		return dep.resU.Reject(report.Message, false)
	} else {
		return nil
	}
}
