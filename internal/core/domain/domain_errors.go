package domain

import "toktok-backend-v1.0.1/pkg/errors"

var (
	ErrInternalServerError  = errors.New("internal server errors")          // 500 Internal Server Error
	ErrNotFound             = errors.New("item is not found")               // 404 Not Found
	ErrConflict             = errors.New("item is conflict")                // 409 Conflict
	ErrBadParam             = errors.New("given param is not valid")        // 400 Bad Request
	ErrDatabaseOperation    = errors.New("database operation failed")       // 500 Internal Server Error
	ErrInvalidTransaction   = errors.New("invalid transaction")             // 400 Bad Request
	ErrForeignKeyViolation  = errors.New("foreign key constraint violated") // 409 Conflict
	ErrUniqueKeyViolation   = errors.New("unique key constraint violated")  // 409 Conflict
	ErrCheckConstraint      = errors.New("check constraint violated")       // 422 Unprocessable Entity
	ErrInvalidQuery         = errors.New("invalid query")                   // 400 Bad Request
	ErrUnsupportedOperation = errors.New("unsupported operation")           // 501 Not Implemented
)
