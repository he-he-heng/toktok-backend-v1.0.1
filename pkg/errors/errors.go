package errors

import (
	stderrors "errors"
)

// New 함수는 text의 새로운 에러를 생성합니다.
func New(text string) error {
	return stderrors.New(text)
}

// Is 함수는 err가 target이 값은지 확인하여 참 거짓을 반환합니다. (standart error와 같습니다)
func Is(err, target error) bool { return stderrors.Is(err, target) }

// As 함수는 err가 target으로 유형이 같은지 확인하여 같다몉 참 혹은 거짓과 함께 변환한 값을 제공합니다.
func As(err error, target interface{}) bool { return stderrors.As(err, target) }

// withMessage 구조체는 에러의 형태를 구조적으로 저장합니다.
type withMessage struct {
	cause error
	msg   string
}

// Wrap 함수는 cause와 help를 인자로하여 warp 하여 에러를 반환합니다.
func Wrap(cause error, help any) error {
	var msg string

	switch v := help.(type) {
	case string:
		msg = v
	case error:
		msg = v.Error()
	default:
		panic("only string and error")
	}

	return &withMessage{
		cause: cause,
		msg:   msg,
	}
}

// Error 함수는 msg와 cuase를 합친 문자열을 반환하는 Error 인터페이스를 구현하는 주체입니다.
func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }

// Cause 함수는 withMessage구조체의 cause를 반환합니다.
func (w *withMessage) Cause() error { return w.cause }

// Unwrap 함수는 Cause 함수와 똑같이 동작하며 withMessage 구조체의 cuase를 반환합니다.
func (w *withMessage) Unwrap() error { return w.cause }

// Cause 함수는 err 인자 값의 시초의 error값을 찾은 후 error를 반환합니다.
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

// Equal 함수는 해당 에러의 시초가 target과 같은 에러인지 비교하여 참, 거짓을 반환합니다.
// Eample
//
//	if Equal(err, domain.InvalidTokenErr) {
//	    fmt.Println("It's Equal!")
//	}
func DeepEqual(err error, target error) bool {
	return Is(Cause(err), target)
}
