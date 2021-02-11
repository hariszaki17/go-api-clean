package exception

// ValidationError expose global
type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}