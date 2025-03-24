package services

import (
	"cloudflare-challenge-weaher-api/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func RedisCacheData(request echo.Context, weatherData map[string]interface{}, err error) error {
	weatherDataJSON, err := json.Marshal(weatherData)
	if err != nil {
		return request.JSON(http.StatusInternalServerError, map[string]string{
			utils.Error: utils.ErrorFailMarshalWD,
		})
	}

	ctx, redisClient := InitConnection()

	err = redisClient.Set(ctx, utils.Weather, weatherDataJSON, 10*time.Minute).Err()
	if err != nil {
		fmt.Println(utils.ErrorRedisSetErr, err)
		return request.JSON(http.StatusInternalServerError, map[string]string{
			utils.Error: utils.ErrorFailCacheData + err.Error(),
		})
	}

	cachedData, err := redisClient.Get(ctx, "weather").Result()
	if err != nil {
		fmt.Println(utils.ErrorRedisGetErr, err)
		return request.JSON(http.StatusInternalServerError, map[string]string{
			utils.Error: utils.ErrorFailRetrCacheData + err.Error(),
		})
	}

	var cachedWeatherData map[string]interface{}
	if err := json.Unmarshal([]byte(cachedData), &cachedWeatherData); err != nil {
		return request.JSON(http.StatusInternalServerError, map[string]string{
			utils.Error: utils.ErrorFailUnMarshalCD,
		})
	}

	return request.JSON(http.StatusOK, cachedWeatherData)
}
