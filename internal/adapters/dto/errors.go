package dto

type APIError struct {
	HttpCode int
	PgCode   string
	Message  string
}

func (e *APIError) Error() string {
	return e.Message
}
