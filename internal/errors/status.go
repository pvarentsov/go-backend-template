package errors

type Status string

const (
	BadRequestError       Status = "BadRequestError"
	InternalError         Status = "InternalError"
	ValidationError       Status = "ValidationError"
	DatabaseError         Status = "DatabaseError"
	NotFoundError         Status = "NotFoundError"
	AlreadyExistsError    Status = "AlreadyExistsError"
	WrongCredentialsError Status = "WrongCredentialsError"
	UnauthorizedError     Status = "UnauthorizedError"
)

func (s Status) Message() string {
	switch s {
	case BadRequestError:
		return "bad request error"
	case InternalError:
		return "internal error"
	case ValidationError:
		return "validation error"
	case DatabaseError:
		return "database error"
	case NotFoundError:
		return "not found error"
	case AlreadyExistsError:
		return "already exists error"
	case WrongCredentialsError:
		return "wrong credentials error"
	case UnauthorizedError:
		return "unauthorized error"
	default:
		return "internal error error"
	}
}
