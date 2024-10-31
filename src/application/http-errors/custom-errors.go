package http_errors

type HttpError struct {
	StatusCode int
	Message string
}

func NewHttpError(statusCode int, message string) *HttpError {
	return &HttpError{
		StatusCode: statusCode,
		Message: message,
	}
}