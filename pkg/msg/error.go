package msg

import (
	"fmt"
	"net/http"
)

var (
	UnauthorizedError = NewError(10001, "Authentication %s %s failed could not found the X-UID field ", 2, nil)
	R_NOT_STARTED     = NewError(10002, "Recruitment %s has not started yet", 1, nil)
	R_ENDED           = NewError(10003, "Recruitment %s has already ended", 1, nil)
	R_ENDED_LONG      = NewError(10004, "Recruitment %s has already ended, hence you cannot modify it. If you REALLY want to extend the end date of this recruitment, please contact maintainers. This is not a bug.", 1, nil)
)

type Error struct {
	id       int
	msg      string
	paramNum int
	details  []string
}

func NewError(id int, msg string, paramNum int, details []string) *Error {
	resp := &Error{
		id:       id,
		msg:      msg,
		paramNum: paramNum,
		details:  details,
	}
	return resp
}
func (resp *Error) StatusCode() int {
	switch resp.id {

	//fail
	case R_NOT_STARTED.id:
		return http.StatusForbidden
	case R_ENDED.id:
		return http.StatusForbidden
	case R_ENDED_LONG.id:
		return http.StatusForbidden
	}
	return http.StatusInternalServerError
}
func (resp *Error) Msg() string {
	return resp.msg
}
func (resp *Error) Details() []string {
	return resp.details
}
func (resp *Error) WithData(data ...string) *Error {
	if len(data) != resp.paramNum {
		return resp
	}
	return &Error{
		id:       resp.id,
		paramNum: resp.paramNum,
		msg:      fmt.Sprintf(resp.msg, data),
		details:  resp.details,
	}
}
func (resp *Error) WithDetail(data ...string) *Error {
	return &Error{
		id:       resp.id,
		paramNum: resp.paramNum,
		msg:      fmt.Sprintf(resp.msg, data),
		details:  append(resp.details, data...),
	}
}

/*
   R_NOT_STARTED: (name: string) => `Recruitment ${name} has not started yet`,
   R_ENDED: (name: string) => `Recruitment ${name} has already ended`,
   // eslint-disable-next-line max-len
   R_ENDED_LONG: (name: string) => `Recruitment ${name} has already ended, hence you cannot modify it. If you REALLY want to extend the end date of this recruitment, please contact maintainers. This is not a bug.`,

*/
