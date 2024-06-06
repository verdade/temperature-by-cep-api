package temperature

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/verdade/temperature-by-cep-api/internal/entity"
	"github.com/verdade/temperature-by-cep-api/pkg/address"
	mock_address "github.com/verdade/temperature-by-cep-api/pkg/address/mock"
	"github.com/verdade/temperature-by-cep-api/pkg/temperature"
	mock_temperature "github.com/verdade/temperature-by-cep-api/pkg/temperature/mock"

	"go.uber.org/mock/gomock"
)

type FindTemperatureTestSuite struct {
	suite.Suite
	AddressInfo            *address.Info
	TemperatureInfo        *temperature.Info
	AddressFetcherMock     *mock_address.MockAddressFetcher
	TemperatureFetcherMock *mock_temperature.MockTemperatureFetcher
}

func (suite *FindTemperatureTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())

	suite.AddressInfo = &address.Info{City: "Imperatriz"}
	suite.TemperatureInfo = &temperature.Info{Celsius: 20, Kelvin: 40, Fahrenheit: 60}
	suite.AddressFetcherMock = mock_address.NewMockAddressFetcher(ctrl)
	suite.TemperatureFetcherMock = mock_temperature.NewMockTemperatureFetcher(ctrl)
}

func (suite *FindTemperatureTestSuite) TestExecute() {
	suite.Run("should return temperature info with successfully", func() {
		suite.AddressFetcherMock.EXPECT().
			GetByZipCode(gomock.Any(), gomock.Any()).
			Return(&address.Info{City: "Imperatriz"}, nil)

		suite.TemperatureFetcherMock.EXPECT().
			GetByCity(gomock.Any(), gomock.Any()).
			Return(&temperature.Info{Celsius: 20, Kelvin: 40, Fahrenheit: 60}, nil)

		findTempUseCase := NewFindTemperatureUseCase(suite.AddressFetcherMock, suite.TemperatureFetcherMock)

		ctx := context.Background()
		tempInfo, err := findTempUseCase.Execute(ctx, "12345678")

		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), tempInfo)
		assert.Equal(suite.T(), suite.TemperatureInfo, tempInfo)
	})

	suite.Run("should return error when address api return error", func() {
		suite.AddressFetcherMock.EXPECT().
			GetByZipCode(gomock.Any(), gomock.Any()).
			Return(nil, entity.ErrZipCodeNotFound)

		findTempUseCase := NewFindTemperatureUseCase(suite.AddressFetcherMock, suite.TemperatureFetcherMock)

		ctx := context.Background()
		tempInfo, err := findTempUseCase.Execute(ctx, "12345678")

		assert.NotNil(suite.T(), err)
		assert.Nil(suite.T(), tempInfo)
	})

	suite.Run("should return error when temperature api return error", func() {
		suite.AddressFetcherMock.EXPECT().
			GetByZipCode(gomock.Any(), gomock.Any()).
			Return(&address.Info{City: "Imperatriz"}, nil)

		suite.TemperatureFetcherMock.EXPECT().
			GetByCity(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("error"))

		findTempUseCase := NewFindTemperatureUseCase(suite.AddressFetcherMock, suite.TemperatureFetcherMock)

		ctx := context.Background()
		tempInfo, err := findTempUseCase.Execute(ctx, "12345678")

		assert.NotNil(suite.T(), err)
		assert.Nil(suite.T(), tempInfo)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(FindTemperatureTestSuite))
}
