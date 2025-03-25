package services_test

import (
	"context"
	"testing"

	"cloudflare-challenge-weaher-api/internal/services"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Ping(ctx context.Context) *redis.StatusCmd {
	args := m.Called(ctx)
	return args.Get(0).(*redis.StatusCmd)
}

func (m *MockRedisClient) Options() *redis.Options {
	args := m.Called()
	return args.Get(0).(*redis.Options)
}

func TestInitConnection(t *testing.T) {

	tests := []struct {
		name         string
		host         string
		pingResult   string
		pingError    error
		expectError  bool
		expectOutput string
	}{
		{
			name:         "Successful connection",
			host:         "localhost:6379",
			pingResult:   "PONG",
			pingError:    nil,
			expectError:  false,
			expectOutput: "Redis connection successful",
		},
		{
			name:         "Failed connection",
			host:         "invalid:6379",
			pingResult:   "",
			pingError:    assert.AnError,
			expectError:  true,
			expectOutput: "Failed to connect to Redis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			originalNewClient := redisNewClient
			defer func() { redisNewClient = originalNewClient }()

			mockClient := new(MockRedisClient)
			redisNewClient = func(opt *redis.Options) *redis.Client {
				assert.Equal(t, tt.host, opt.Addr)
				return &redis.Client{}
			}

			// Mock Ping behavior
			cmd := &redis.StatusCmd{}
			if tt.pingError != nil {
				cmd.SetErr(tt.pingError)
			} else {
				cmd.SetVal(tt.pingResult)
			}
			mockClient.On("Ping", mock.Anything).Return(cmd)

			ctx, client := services.InitConnection()

			assert.NotNil(t, ctx)

			assert.NotNil(t, client)
		})
	}
}

var redisNewClient = func(opt *redis.Options) *redis.Client {
	return redis.NewClient(opt)
}
