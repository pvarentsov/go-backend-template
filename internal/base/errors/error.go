package errors

import "fmt"

type Error struct {
	status  Status
	message string
	details string
}

func (e Error) Error() string {
	return e.message
}

func (e Error) ErrorWithDetails() string {
	if e.details != "" {
		return fmt.Sprintf("%s: %s", e.message, e.details)
	}
	return e.message
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

func Wrap(err error, status Status, message string) Error {
	newErr := Error{
		status:  status,
		message: message,
		details: err.Error(),
	}
	if baseErr, ok := err.(Error); ok {
		newErr.details = baseErr.ErrorWithDetails()
	}
	if len(message) == 0 {
		newErr.message = status.Message()
	}

	return newErr
}

func Errorf(status Status, message string, a ...interface{}) Error {
	err := Error{
		status:  status,
		message: fmt.Sprintf(message, a...),
	}
	if len(message) == 0 {
		err.message = status.Message()
	}

	return err
}

func Wrapf(err error, status Status, message string, a ...interface{}) Error {
	newErr := Error{
		status:  status,
		message: fmt.Sprintf(message, a...),
		details: err.Error(),
	}
	if len(message) == 0 {
		newErr.message = status.Message()
	}
	if baseErr, ok := err.(Error); ok {
		newErr.details = baseErr.ErrorWithDetails()
	}

	return newErr
}
