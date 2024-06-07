package handler

import (
	"encoding/json"
	"net/http"

	"github.com/verdade/temperature-by-cep-api/internal/entity"
	"github.com/verdade/temperature-by-cep-api/pkg/logger"
	"github.com/verdade/temperature-by-cep-api/pkg/temperature"
)

type WebTemperatureHandler struct {
	FindTemperature entity.FindTemperatureUseCase
}

type TemperaturePresenter struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func NewWebTemperatureHandler(findTemperature entity.FindTemperatureUseCase) *WebTemperatureHandler {
	return &WebTemperatureHandler{
		FindTemperature: findTemperature,
	}
}

func (h *WebTemperatureHandler) TemperatureByCepHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	logger.Info("[TemperatureByCepHandler] starting handler")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	zipcode := r.URL.Query().Get("cep")

	if !zipCodeValidator(zipcode) {
		logger.Error("[TemperatureByCepHandler] invalid zipcode", nil)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}

	tempData, err := h.FindTemperature.Execute(ctx, zipcode)

	if err == entity.ErrZipCodeNotFound {
		logger.Error("[TemperatureByCepHandler] zipcode not found", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not found zipcode"))
		return
	}

	if err != nil {
		logger.Error("[TemperatureByCepHandler] fail to execute usecase FindTemperature", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	tempDataPresenter := toTemperaturePresenter(tempData)
	json.NewEncoder(w).Encode(&tempDataPresenter)
}

func toTemperaturePresenter(temp *temperature.Info) *TemperaturePresenter {
	return &TemperaturePresenter{
		TempC: temp.Celsius,
		TempF: temp.Fahrenheit,
		TempK: temp.Kelvin,
	}
}

func zipCodeValidator(zipCode string) bool {
	return len(zipCode) == 8
}
