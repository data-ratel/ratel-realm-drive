package error

type BaseError struct {
	Error string `json:"error"`
}

// HTTP error result wrapper
type ErrorResult struct {
	Error string `json:"error"`
}
