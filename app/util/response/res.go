package respU
import "github.com/og/x/error"

func Reject(msg string, shouldRecord bool) error {
	return ge.NewReject(Response{
		Type: Response{}.Type.Enum().Fail,
		Msg: msg,
	}, shouldRecord)
}
type Response struct {
	Type ResponseType `json:"type"`
	Msg string `json:"msg"`
	FailCode string `json:"failCode"`
	Data interface{} `json:"data"`
}
type ResponseType string
func (v ResponseType) Enum() (e struct {
	Pass ResponseType
	Fail ResponseType
}) {
	e.Pass = "pass"
	e.Fail = "fail"
	return
}