package entity

import (
	"context"
	"errors"

	"github.com/verdade/temperature-by-cep-api/pkg/temperature"
)

var ErrZipCodeNotFound error = errors.New("zipcode not found")

type FindTemperatureUseCase interface {
	Execute(ctx context.Context, zipCode string) (*temperature.Info, error)
}
