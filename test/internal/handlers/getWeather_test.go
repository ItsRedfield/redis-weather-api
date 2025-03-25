package handlers_test

import (
	"cloudflare-challenge-weaher-api/pkg/utils"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockServices struct {
	mock.Mock
}

func (m *MockServices) FetchWeatherData(lat, lon string) (interface{}, error) {
	args := m.Called(lat, lon)
	return args.Get(0), args.Error(1)
}

func (m *MockServices) RedisCacheData(request echo.Context, data interface{}, err error) error {
	args := m.Called(request, data, err)
	return args.Error(0)
}

func TestGetWeather(t *testing.T) {
	e := echo.New()
	mockServices := new(MockServices)

	tests := []struct {
		name           string
		lat            string
		lon            string
		fetchDataResp  interface{}
		fetchDataErr   error
		cacheDataErr   error
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name:           "Valid coordinates",
			lat:            "40.7128",
			lon:            "-74.0060",
			fetchDataResp:  map[string]interface{}{"temp": "20"},
			fetchDataErr:   nil,
			cacheDataErr:   nil,
			expectedStatus: http.StatusOK,
			expectedBody:   nil,
		},
		{
			name:           "Invalid latitude",
			lat:            "invalid",
			lon:            "-74.0060",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]string{utils.Error: utils.BadLatOrLonMsg},
		},
		{
			name:           "Invalid longitude",
			lat:            "40.7128",
			lon:            "invalid",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]string{utils.Error: utils.BadLatOrLonMsg},
		},
		{
			name:           "FetchWeatherData error",
			lat:            "40.7128",
			lon:            "-74.0060",
			fetchDataErr:   errors.New(utils.ErrorInvalidCoord),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]string{utils.Error: utils.ErrorInvalidCoord},
		},
		{
			name:           "RedisCacheData error",
			lat:            "40.7128",
			lon:            "-74.0060",
			fetchDataResp:  map[string]interface{}{"temp": "20"},
			cacheDataErr:   errors.New("cache error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   map[string]string{utils.Error: "cache error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/weather", nil)
			q := req.URL.Query()
			q.Add(utils.LatQueryParam, tt.lat)
			q.Add(utils.LonQueryParam, tt.lon)
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if tt.fetchDataErr != nil || tt.fetchDataResp != nil {
				mockServices.On("FetchWeatherData", tt.lat, tt.lon).Return(tt.fetchDataResp, tt.fetchDataErr)
			}
			if tt.cacheDataErr != nil {
				mockServices.On("RedisCacheData", c, tt.fetchDataResp, tt.fetchDataErr).Return(tt.cacheDataErr)
			}

		})
	}
}
