package message

import "errors"

const (
	Success = iota + 1
	Fail
)

const (
	LIMIT_RANGE_MSG = 1
)


var NotImplementError = errors.New("have not implemented")
var errorText = map[int32]string{
	Success: "Thành công",
	Fail:    "Có lỗi xảy ra khi xử lí yêu cầu",
}

func ErrorMessage(code int32) string {
	return errorText[code]
}
