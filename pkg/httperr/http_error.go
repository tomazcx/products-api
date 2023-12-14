package httperr

type HttpError struct {
	StatusCode int
	Message    string
}

func NewHttpError(message string, statusCode int) *HttpError {
	return &HttpError{Message: message, StatusCode: statusCode}
}

func (e *HttpError) Error() string {
	return e.Message
}
