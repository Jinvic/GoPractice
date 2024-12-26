package e

type Err struct {
	code int
	msg  string
}

func New(code int, msg string) Err {
	return Err{
		code: code,
		msg:  msg,
	}
}

func (e Err) Error() string {
	return e.msg
}

var (
	DATA_NOT_EXIST = Err{10001, "data not exist"}
)
