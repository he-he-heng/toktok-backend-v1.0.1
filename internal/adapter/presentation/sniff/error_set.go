package sniff

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"toktok-backend-v1.0.1/internal/core/domain"
	"toktok-backend-v1.0.1/pkg/errors"
)

var es *errorSet
var once sync.Once

type statusAndErr struct {
	status int
	err    error
}

func newStatusAndErr(status int, err error) statusAndErr {
	return statusAndErr{
		status: status,
		err:    err,
	}
}

type errorSet struct {
	statusAndErrs []statusAndErr
}

func (es *errorSet) add(statusCode int, errs ...error) {
	for _, err := range errs {
		es.statusAndErrs = append(es.statusAndErrs, newStatusAndErr(statusCode, err))
	}
}

func (es *errorSet) Get(argErr error) (int, string, error) {
	for _, statusAndErr := range es.statusAndErrs {
		if errors.DeepEqual(argErr, statusAndErr.err) {
			return statusAndErr.status, statusAndErr.err.Error(), nil
		}
	}

	return 0, "", errors.New("err not add in errorSet")
}

func errSetGet() *errorSet {
	if es == nil {
		once.Do(func() {
			es = &errorSet{
				statusAndErrs: make([]statusAndErr, 10),
			}
		})

		// StatusBadRequest 400
		es.add(fiber.StatusBadRequest, domain.ErrBadParam, domain.ErrInvalidTransaction, domain.ErrInvalidQuery)

		// StatusUnauthorized 401
		es.add(fiber.StatusUnauthorized, domain.ErrUnauthorized)

		// StatusNotFound 404
		es.add(fiber.StatusNotFound, domain.ErrNotFound)

		// StatusConflict 409
		es.add(fiber.StatusConflict, domain.ErrConflict, domain.ErrForeignKeyViolation, domain.ErrUniqueKeyViolation)

		// StatusUnprocessableEntity 422
		es.add(fiber.StatusUnprocessableEntity, domain.ErrCheckConstraint)

		// StatusInternalServerError 500
		es.add(fiber.StatusInternalServerError, domain.ErrDatabaseOperation, domain.ErrInternalServerError)

		// StatusNotImplemented 501
		es.add(fiber.StatusNotImplemented, domain.ErrUnsupportedOperation)
	}

	return es
}
