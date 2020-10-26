package vdU

import (
	respU "github.com/og/goclub/app/util/response"
	vd "github.com/og/juice/validator"
)

var checker vd.Checker
func init () {
	checker = vd.NewCN()
}
func Check(data vd.Data) error {
	report := checker.Check(data)
	if report.Fail {
		return respU.Reject(report.Message, false)
	} else {
		return nil
	}
}
