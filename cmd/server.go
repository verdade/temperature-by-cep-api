package main

import (
	"github.com/verdade/temperature-by-cep-api/configs"
	"github.com/verdade/temperature-by-cep-api/internal/infra/web/handler"
	"github.com/verdade/temperature-by-cep-api/internal/infra/web/webserver"
	"github.com/verdade/temperature-by-cep-api/internal/usecase/temperature"
	"github.com/verdade/temperature-by-cep-api/pkg/address/viacep"
	"github.com/verdade/temperature-by-cep-api/pkg/requester/resty"
	"github.com/verdade/temperature-by-cep-api/pkg/temperature/weather"
)

func main() {
	configs, err := configs.LoadConfig(".temperature-by-cep-api")
	if err != nil {
		panic(err)
	}

	requester := resty.New()
	viaCepFetcher := viacep.New(configs.ViaCepApiUrl, requester)
	weatherFetcher := weather.New(configs.WeatherApiUrl, configs.WeatherApiKey, requester)

	findTemperature := temperature.NewFindTemperatureUseCase(viaCepFetcher, weatherFetcher)

	ws := webserver.New(configs.WebServerPort)
	tH := handler.NewWebTemperatureHandler(findTemperature)
	ws.AddHandler("/temperature", tH.TemperatureByCepHandler)
	ws.Start()
}
