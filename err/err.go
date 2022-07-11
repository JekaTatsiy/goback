package err

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

type errCode int

type Err struct {
	innerError bool
	Code       errCode
	Msg        string
	Func       string
	Line       int
}

func NewErr(code errCode, msg string) *Err {
	return &Err{Code: code, Msg: msg}
}

func FromMsg(msg string) *Err {
	return &Err{Code: E_SOME, Msg: msg}
}

func FromMsgf(format string, a ...any) *Err {
	return &Err{Code: E_SOME, Msg: fmt.Sprintf(format, a...)}
}

func FromCode(code errCode) *Err {
	msg, ok := msgs[code]
	if !ok {
		return &Err{innerError: true, Code: 1, Msg: fmt.Sprintf("non-existent-code: \"%d\"", code)}
	}
	return &Err{Code: code, Msg: msg}
}
func FromError(e error) *Err {
	return &Err{Code: E_SOME, Msg: e.Error()}
}

func NewErrEmp() *Err {
	return &Err{Code: 0, Msg: ""}
}

func (e *Err) WithCode(code int) *Err {
	if e.innerError {
		return e
	}
	e.Code = errCode(code)
	return e
}

func (e *Err) WithMsg(msg string) *Err {
	if e.innerError {
		return e
	}
	e.Msg = msg
	return e
}

func (e *Err) WithPos(dept int) *Err {
	if e.innerError {
		return e
	}
	_, fn, line, ok := runtime.Caller(dept)
	if ok {
		e.Func = fn
		e.Line = line
	}

	return e
}

func (e *Err) Err(l *logrus.Entry) *Err {
	e.WithPos(2)
	l.
		WithField("Code", e.Code).
		WithField("Msg", e.Msg).
		WithField("Func", e.Func).
		WithField("Line", e.Line).
		Error()
	return e
}
func (e *Err) Info(l *logrus.Entry) *Err {
	l.
		WithField("Code", e.Code).
		WithField("Msg", e.Msg).
		WithField("Func", e.Func).
		WithField("Line", e.Line).
		Info()
	return e
}

func (e *Err) Recover() *Err {
	e.innerError = false
	return e
}
