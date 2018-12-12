package message

import (
	"errors"
	"encoding/json"
)

const (
	Success = iota + 1
	Fail
)

const (
	MYSQL = "mysql"
	MEMORY = "memory"
)

const (
	LIMIT_RANGE_MSG = 1000
)


var NotImplementError = errors.New("have not implemented")
var errorText = map[int32]string{
	Success: "operation success",
	Fail:    "fail to process request",
}

func ErrorMessage(code int32) string {
	return errorText[code]
}


type TsRange struct {
	Timestamp int64 `json:"timestamp"`
}

func NewTsRange(ts int64) TsRange  {
	return TsRange{
		Timestamp: ts,
	}
}

func (t TsRange) ToJson() string  {
	b, _ := json.Marshal(&t)
	return string(b)
}