package core

import "fmt"

// ErrInvalidField indicates that a field provided by the client is invalid
type ErrInvalidField struct {
	Field string
}

func (e *ErrInvalidField) Error() string {
	return fmt.Sprintf("Invalid field '%s'", e.Field)
}
