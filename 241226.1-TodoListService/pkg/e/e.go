package e

type Err struct {
	Code int
	Msg  string
}

func New(code int, msg string) Err {
	return Err{
		Code: code,
		Msg:  msg,
	}
}

func (e Err) Error() string {
	return e.Msg
}

var (
	DATA_NOT_EXIST = Err{10001, "数据不存在"}
	PARAM_ERROR   = Err{10002, "参数错误"}
)
