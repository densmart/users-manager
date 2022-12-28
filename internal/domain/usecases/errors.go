package usecases

type APIError struct {
	HttpCode int
	Message  string
}

func (e *APIError) Error() string {
	return e.Message
}
