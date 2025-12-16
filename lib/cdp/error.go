package cdp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Error of the Response.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// UnmarshalJSON implements json.Unmarshaler to handle error codes as either int or string.
func (e *Error) UnmarshalJSON(data []byte) error {
	// decode into an auxiliary struct to ambiguously typed Code field
	type errorAlias Error
	aux := &struct {
		Code json.RawMessage `json:"code"`
		*errorAlias
	}{
		errorAlias: (*errorAlias)(e),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// check whether the code field has at least one character
	if len(aux.Code) > 0 {
		var intCode int
		if err := json.Unmarshal(aux.Code, &intCode); err == nil {
			e.Code = intCode
		} else {
			var stringCode string
			if err := json.Unmarshal(aux.Code, &stringCode); err == nil {
				if parsed, err := strconv.Atoi(stringCode); err == nil {
					e.Code = parsed
				}
			}
		}
	}

	return nil
}

// Error stdlib interface.
func (e *Error) Error() string {
	return fmt.Sprintf("%v", *e)
}

// Is stdlib interface.
func (e Error) Is(target error) bool {
	err, ok := target.(*Error)
	return ok && e == *err
}

// ErrCtxNotFound type.
var ErrCtxNotFound = &Error{
	Code:    -32000,
	Message: "Cannot find context with specified id",
}

// ErrSessionNotFound type.
var ErrSessionNotFound = &Error{
	Code:    -32001,
	Message: "Session with given id not found.",
}

// ErrSearchSessionNotFound type.
var ErrSearchSessionNotFound = &Error{
	Code:    -32000,
	Message: "No search session with given id found",
}

// ErrCtxDestroyed type.
var ErrCtxDestroyed = &Error{
	Code:    -32000,
	Message: "Execution context was destroyed.",
}

// ErrObjNotFound type.
var ErrObjNotFound = &Error{
	Code:    -32000,
	Message: "Could not find object with given id",
}

// ErrNodeNotFoundAtPos type.
var ErrNodeNotFoundAtPos = &Error{
	Code:    -32000,
	Message: "No node found at given location",
}

// ErrNotAttachedToActivePage type.
var ErrNotAttachedToActivePage = &Error{
	Code:    -32000,
	Message: "Not attached to an active page",
}
