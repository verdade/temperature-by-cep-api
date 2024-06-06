package temperature

import (
	"context"

	"github.com/verdade/temperature-by-cep-api/internal/entity"
	"github.com/verdade/temperature-by-cep-api/pkg/address"
	"github.com/verdade/temperature-by-cep-api/pkg/logger"
	"github.com/verdade/temperature-by-cep-api/pkg/temperature"
)

type FindTemperatureUseCase struct {
	AddressFetcher     address.AddressFetcher
	TemperatureFetcher temperature.TemperatureFetcher
}

func NewFindTemperatureUseCase(af address.AddressFetcher, tf temperature.TemperatureFetcher) *FindTemperatureUseCase {
	return &FindTemperatureUseCase{
		AddressFetcher:     af,
		TemperatureFetcher: tf,
	}
}

func (s *FindTemperatureUseCase) Execute(ctx context.Context, zipCode string) (*temperature.Info, error) {
	logger.Info("[FindTemperatureUseCase] starting usecase for zipcode: " + zipCode)
	address, err := s.AddressFetcher.GetByZipCode(ctx, zipCode)
	if err != nil {
		logger.Error("[FindTemperatureUseCase] fail to execute AddressApi for zipcode: "+zipCode, err)
		return nil, entity.ErrZipCodeNotFound
	}
	temp, err := s.TemperatureFetcher.GetByCity(ctx, address.City)
	if err != nil {
		logger.Error("[FindTemperatureUseCase] fail to execute TemperatureApi  for zipcode: "+zipCode, err)
		return nil, err
	}
	logger.Info("[FindTemperatureUseCase] finishing usecase for zipcode: " + zipCode + " with success")
	return temp, nil
}
