package utils

import (
	"toktok-backend-v1.0.1/pkg/errors"

	"gorm.io/gorm"
	"toktok-backend-v1.0.1/internal/core/domain"
)

// WrapErr 함수는 gorm.Err를 인자값으로 대입하면 domain.error 값을 반환하는 함수입니다.
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return domain.ErrNotFound
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return domain.ErrUniqueKeyViolation
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return domain.ErrForeignKeyViolation
	case errors.Is(err, gorm.ErrCheckConstraintViolated):
		return domain.ErrCheckConstraint
	case errors.Is(err, gorm.ErrInvalidTransaction):
		return domain.ErrInvalidTransaction
	case errors.Is(err, gorm.ErrInvalidData),
		errors.Is(err, gorm.ErrInvalidField),
		errors.Is(err, gorm.ErrInvalidValue),
		errors.Is(err, gorm.ErrInvalidValueOfLength):
		return domain.ErrBadParam
	case errors.Is(err, gorm.ErrMissingWhereClause),
		errors.Is(err, gorm.ErrPrimaryKeyRequired),
		errors.Is(err, gorm.ErrModelValueRequired),
		errors.Is(err, gorm.ErrModelAccessibleFieldsRequired),
		errors.Is(err, gorm.ErrSubQueryRequired):
		return domain.ErrInvalidQuery
	case errors.Is(err, gorm.ErrNotImplemented),
		errors.Is(err, gorm.ErrUnsupportedRelation),
		errors.Is(err, gorm.ErrUnsupportedDriver),
		errors.Is(err, gorm.ErrDryRunModeUnsupported):
		return domain.ErrUnsupportedOperation
	case errors.Is(err, gorm.ErrRegistered),
		errors.Is(err, gorm.ErrEmptySlice),
		errors.Is(err, gorm.ErrInvalidDB),
		errors.Is(err, gorm.ErrPreloadNotAllowed):
		return domain.ErrDatabaseOperation
	default:
		return domain.ErrInternalServerError
	}
}
