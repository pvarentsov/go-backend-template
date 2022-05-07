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

func Wrap(status Status, err error, message string) Error {
	coreErr, isCoreErr := err.(Error)

	if isCoreErr {
		return coreErr
	}

	return New(status, message).setInternal(err)
}

func Wrapf(status Status, err error, message string, a ...interface{}) Error {
	coreErr, isCoreErr := err.(Error)

	if isCoreErr {
		return coreErr
	}

	return Errorf(status, message, a...).setInternal(err)
}

func (e Error) setInternal(internal error) Error {
	e.internal = internal
	return e
}
