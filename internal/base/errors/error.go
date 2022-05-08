package errors

import "fmt"

type Error struct {
	status  Status
	message string
	err     error
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) DetailedError() string {
	if e.err != nil {
		if baseErr, ok := e.err.(*Error); ok {
			return fmt.Sprintf("%s: %s", e.message, baseErr.DetailedError())
		}
		return fmt.Sprintf("%s: %s", e.message, e.err.Error())
	}
	return e.message
}

func (e *Error) Status() Status {
	return e.status
}

func (e *Error) Unwrap() error {
	return e.err
}

func New(status Status, message string) *Error {
	err := Error{
		status:  status,
		message: message,
	}
	if len(message) == 0 {
		err.message = status.Message()
	}

	return &err
}

func Errorf(status Status, message string, a ...interface{}) *Error {
	err := Error{
		status:  status,
		message: fmt.Sprintf(message, a...),
	}
	if len(message) == 0 {
		err.message = status.Message()
	}

	return &err
}

func Wrap(err error, status Status, message string) *Error {
	newErr := Error{
		status:  status,
		message: message,
		err:     err,
	}
	if len(message) == 0 {
		newErr.message = status.Message()
	}

	return &newErr
}

func Wrapf(err error, status Status, message string, a ...interface{}) *Error {
	newErr := Error{
		status:  status,
		message: fmt.Sprintf(message, a...),
		err:     err,
	}
	if len(message) == 0 {
		newErr.message = status.Message()
	}

	return &newErr
}
