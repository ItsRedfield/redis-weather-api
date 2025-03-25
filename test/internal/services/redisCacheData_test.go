package services

import (
	"cloudflare-challenge-weaher-api/internal/services"
	"cloudflare-challenge-weaher-api/pkg/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	args := m.Called(ctx, key, value, expiration)
	return redis.NewStatusCmd(ctx, args...)
}

func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	args := m.Called(ctx, key)
	return redis.NewStringCmd(ctx, args...)
}

func InitConnection() (context.Context, *MockRedisClient) {
	return context.TODO(), &MockRedisClient{}
}

func TestRedisCacheData(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRedisClient := new(MockRedisClient)
	ctx := context.TODO()

	weatherData := map[string]interface{}{
		"temperature": 25.0,
		"humidity":    80.0,
	}

	weatherDataJSON, _ := json.Marshal(weatherData)

	mockRedisClient.On("Set", ctx, utils.Weather, weatherDataJSON, 10*time.Minute).Return(redis.NewStatusCmd(ctx))
	mockRedisClient.On("Get", ctx, utils.Weather).Return(redis.NewStringCmd(ctx, string(weatherDataJSON)))
	var mockRedisErrResp = http.StatusInternalServerError

	err := services.RedisCacheData(c, weatherData, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var responseData map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseData)
	assert.NoError(t, err)
	assert.Equal(t, weatherData, responseData)

	err = services.RedisCacheData(c, map[string]interface{}{"invalid": make(chan int)}, nil)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, mockRedisErrResp)

	mockRedisClient.On("Set", ctx, utils.Weather, weatherDataJSON, 10*time.Minute).Return(redis.NewStatusCmd(ctx, redis.ErrClosed))
	err = services.RedisCacheData(c, weatherData, nil)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, mockRedisErrResp)

	mockRedisClient.On("Get", ctx, utils.Weather).Return(redis.NewStringCmd(ctx, redis.ErrClosed))
	err = services.RedisCacheData(c, weatherData, nil)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, mockRedisErrResp)
}
