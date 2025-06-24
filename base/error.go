package base

import (
	"bytes"
	"fmt"
	"google.golang.org/grpc/codes"
)

type Error struct {
	code        codes.Code
	event       string
	msg         string
	attachedErr error
}

func NewErr(code codes.Code, event string) *Error {
	return &Error{
		code:  code,
		event: event,
	}
}

func (e *Error) AttachErr(err error) *Error {
	e.attachedErr = err
	return e
}

func (e *Error) SetMsg(format string, args ...any) *Error {
	e.msg = fmt.Sprintf(format, args...)
	return e
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	buf := bytes.NewBuffer(nil)
	_, _ = fmt.Fprintf(buf, "code=%s", e.code.String())
	if e.event != "" {
		_, _ = fmt.Fprintf(buf, ";event=%s", e.event)
	}
	if e.msg != "" {
		_, _ = fmt.Fprintf(buf, ";msg=%s", e.msg)
	}
	if e.attachedErr != nil {
		_, _ = fmt.Fprintf(buf, ";err=%s", e.attachedErr.Error())
	}
	return buf.String()
}
