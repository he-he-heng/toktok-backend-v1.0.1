package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type v struct {
	instance *validator.Validate
}

func newVerifier() *v {
	verifier := v{
		instance: validator.New(validator.WithRequiredStructEnabled()),
	}

	return &verifier
}

func (v v) Validate(s any) error {
	return v.instance.Struct(s)
}

var once sync.Once
var instance *v

func Get() *v {
	once.Do(func() {
		//init
		instance = newVerifier()

		// register custom tag

	})

	return instance
}
