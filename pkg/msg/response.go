package msg

import (
	"net/http"
)

var (
	R_SUCCESS = NewResponseDetails(100001, "recruitment get success!", "")
)

var (
	R_NOT_STARTED = NewResponseDetails(1, "Recruitment %s has not started yet", "")
	R_ENDED       = NewResponseDetails(2, "Recruitment %s has already ended", "")
	R_ENDED_LONG  = NewResponseDetails(3, "Recruitment %s has already ended, hence you cannot modify it. If you REALLY want to extend the end date of this recruitment, please contact maintainers. This is not a bug.", "")
)

type ResponseDetails struct {
	Id   int         `json:"id"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponseDetails(id int, msg string, data interface{}) *ResponseDetails {
	resp := &ResponseDetails{
		Id:   id,
		Msg:  msg,
		Data: data,
	}
	return resp
}
func (resp *ResponseDetails) StatusCode() int {
	switch resp.Id {
	//success
	case R_SUCCESS.Id:
		return http.StatusAccepted

	//fail
	case R_NOT_STARTED.Id:
		return http.StatusForbidden
	case R_ENDED.Id:
		return http.StatusForbidden
	case R_ENDED_LONG.Id:
		return http.StatusForbidden
	}
	return http.StatusInternalServerError
}

/*
   R_NOT_STARTED: (name: string) => `Recruitment ${name} has not started yet`,
   R_ENDED: (name: string) => `Recruitment ${name} has already ended`,
   // eslint-disable-next-line max-len
   R_ENDED_LONG: (name: string) => `Recruitment ${name} has already ended, hence you cannot modify it. If you REALLY want to extend the end date of this recruitment, please contact maintainers. This is not a bug.`,

*/
