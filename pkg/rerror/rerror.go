package rerror

import (
	"fmt"
	"net/http"
)

var (
	UnauthorizedError = NewError(10001, "Authentication failed could not get uid", 0, nil)

	SSOError = NewError(10005, "SSO get UserInfo failed ", 0, nil)

	RecruitmentNotReady      = NewError(10002, "Recruitment %s has not started yet", 1, nil)
	RecruitmentEnd           = NewError(10003, "Recruitment %s has already ended", 1, nil)
	RecruitmentEndDontModify = NewError(10004, "Recruitment %s has already ended, hence you cannot modify it. If you REALLY want to extend the end date of this recruitment, please contact maintainers. This is not a bug.", 1, nil)
	RecruitmentStopped       = NewError(10011, "The application deadline of recruitment %s has already passed", 1, nil)

	SendSMSError        = NewError(10006, "Send sms failed", 0, nil)
	RequestBodyError    = NewError(10007, "Request body rerror", 0, nil)
	RequestParamError   = NewError(10015, "Request param rerror", 0, nil)
	SaveDatabaseError   = NewError(10008, "Save %s rerror", 1, nil)
	UpdateDatabaseError = NewError(10009, "Update %s rerror", 1, nil)
	GetDatabaseError    = NewError(10010, "Get %s rerror", 1, nil)

	UpLoadFileError   = NewError(10012, "%s upload file rerror", 1, nil)
	DownloadFileError = NewError(10013, "%s download fail rerror", 1, nil)

	RoleError = NewError(10014, "%s don`t has role to %s", 2, nil)

	Abandoned                = NewError(10015, "Application of %s has already been abandoned ", 1, nil)
	Rejected                 = NewError(10016, "Application of %s has already been rejected", 1, nil)
	GroupNotMatch            = NewError(10017, "Member' group does not match", 0, nil)
	InterviewUpdateError     = NewError(10018, "Update interview rerror, %s", 1, nil)
	InterviewHasBeenSelected = NewError(10019, "Interview %s have been selected", 1, nil)
	ReselectInterviewError   = NewError(10020, "You don't have the perssion to re-select this application", 0, nil)

	CheckPermissionError = NewError(10021, "Check permission rerror", 0, nil)

	RemoveDatabaseError = NewError(10022, "Remove %s rerror", 1, nil)

	SMSSendError = NewError(10023, "Send sms to %s rerror", 1, nil)

	NoInterviewScheduled = NewError(10024, "No interviews are scheduled for %s", 1, nil)
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
	case RequestBodyError.id:
		return http.StatusForbidden
	case RecruitmentNotReady.id:
		return http.StatusForbidden
	case RecruitmentEnd.id:
		return http.StatusForbidden
	case RecruitmentEndDontModify.id:
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
func (resp *Error) WithData(data ...interface{}) *Error {
	if len(data) != resp.paramNum {
		return resp
	}
	return &Error{
		id:       resp.id,
		paramNum: resp.paramNum,
		msg:      fmt.Sprintf(resp.msg, data...),
		details:  resp.details,
	}
}
func (resp *Error) WithDetail(data ...string) *Error {
	return &Error{
		id:       resp.id,
		paramNum: resp.paramNum,
		msg:      resp.msg,
		details:  append(resp.details, data...),
	}
}

/*
   R_NOT_STARTED: (name: string) => `Recruitment ${name} has not started yet`,
   R_ENDED: (name: string) => `Recruitment ${name} has already ended`,
   // eslint-disable-next-line max-len
   R_ENDED_LONG: (name: string) => `Recruitment ${name} has already ended, hence you cannot modify it. If you REALLY want to extend the end date of this recruitment, please contact maintainers. This is not a bug.`,

*/
