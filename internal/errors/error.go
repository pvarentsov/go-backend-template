package errors

import "fmt"

type Error struct {
	status   Status
	internal error
	message  string
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Internal() error {
	return e.internal
}

func (e Error) Status() Status {
	return e.status
}

func (e Error) SetInternal(internal error) Error {
	e.internal = internal
	return e
}

func New(status Status, message string) Error {
	err := Error{
		status:  status,
		message: message,
	}

	if len(message) == 0 {
		err.message = status.Message()
	}

	return err
}

func Errorf(status Status, message string, a ...interface{}) Error {
	err := Error{status: status}

	if len(message) == 0 {
		err.message = status.Message()
		return err
	}

	err.message = fmt.Sprintf(message, a...)

	return err
}

func Wrap(err error) Error {
	coreErr, isCoreErr := err.(Error)

	if isCoreErr {
		return coreErr
	}

	return New(InternalError, "").SetInternal(err)
}
