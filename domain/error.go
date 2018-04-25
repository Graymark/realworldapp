package domain

import "fmt"

type ErrorType int

const (
	ErrLogical ErrorType = iota
	ErrInterface
	ErrConsistency
	ErrSecurity
	ErrInvalidArg
)

var eTstrings = [5]string{
	"LOGICAL",
	"INTERFACE",
	"CONSISTENCY",
	"SECURITY",
	"INVALID_ARGUMENT",
}

func (eT ErrorType) String() string {
	return eTstrings[int(eT)]
}

type ErrorLevel int

const (
	ErrDebug ErrorLevel = iota
	ErrInfo
	ErrWarn
	ErrError
	ErrFatal
	ErrPanic
)

var eLstrings = [6]string{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
	"FATAL",
	"PANIC",
}

func (eL ErrorLevel) String() string {
	return eLstrings[int(eL)]
}

type Error struct {
	ErrorType
	ErrorLevel
	Message    string
	Additional string
}

func (e Error) Error() string {
	return fmt.Sprintf("[%s:%s] %v || %v", e.ErrorLevel.String(), e.ErrorType.String(), e.Message, e.Additional)
}

func NewError(lvl ErrorLevel, eType ErrorType, mess string, args ...interface{}) *Error {
	e := Error{ErrorType: eType, ErrorLevel: lvl, Message: mess}
	if len(args) > 0 {
		a := fmt.Sprint(args...)
		e.Additional = a
	}
	return &e
}

type FormErrors struct {
	Errors []ErrorForm
}

func (FormErrors) Error() string {
	return ""
}

type ErrorForm struct {
	Name    string
	Message string
}
