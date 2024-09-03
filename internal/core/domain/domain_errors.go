package domain

import "toktok-backend-v1.0.1/pkg/errors"

var (
	ErrBadParam             = errors.New("given param is not valid")        // 400 Bad Request
	ErrInvalidTransaction   = errors.New("invalid transaction")             // 400 Bad Request
	ErrInvalidQuery         = errors.New("invalid query")                   // 400 Bad Request
	ErrUnauthorized         = errors.New("unauthorized")                    // 401 Unauthorized
	ErrNotFound             = errors.New("item is not found")               // 404 Not Found
	ErrConflict             = errors.New("item is conflict")                // 409 Conflict
	ErrForeignKeyViolation  = errors.New("foreign key constraint violated") // 409 Conflict
	ErrUniqueKeyViolation   = errors.New("unique key constraint violated")  // 409 Conflict
	ErrCheckConstraint      = errors.New("check constraint violated")       // 422 Unprocessable Entity
	ErrDatabaseOperation    = errors.New("database operation failed")       // 500 Internal Server Error
	ErrInternalServerError  = errors.New("internal server errors")          // 500 Internal Server Error
	ErrUnsupportedOperation = errors.New("unsupported operation")           // 501 Not Implemented

)
