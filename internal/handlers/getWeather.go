package handlers

import (
	"cloudflare-challenge-weaher-api/internal/services"
	"cloudflare-challenge-weaher-api/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetWeather(request echo.Context) error {
	lat := request.QueryParam(utils.LatQueryParam)
	lon := request.QueryParam(utils.LonQueryParam)

	isValidLat, _ := utils.CoordinatePattern(lat)
	isValidLon, _ := utils.CoordinatePattern(lon)

	if !isValidLat || !isValidLon {
		return request.JSON(http.StatusBadRequest, map[string]string{
			utils.Error: utils.BadLatOrLonMsg,
		})
	}

	weatherData, err := services.FetchWeatherData(lat, lon)
	if err != nil {
		if err.Error() == utils.ErrorInvalidCoord {
			return request.JSON(http.StatusBadRequest, map[string]string{
				utils.Error: err.Error(),
			})
		}
		return request.JSON(http.StatusInternalServerError, map[string]string{
			utils.Error: err.Error(),
		})
	}

	return services.RedisCacheData(request, weatherData, err)
}
