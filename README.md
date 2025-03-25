# Cloudflare Challenge Weather API

This project is a weather API built using the Echo framework in Go. It provides endpoints for user login and fetching weather data based on latitude and longitude. The weather data is fetched from the National Weather Service (NWS) API and cached using Redis.

## Project Structure

The Project have the following structure:

```sh
.
|-- cmd
|   |-- cloudflare-challenge-weather-api
|       |-- main.go
|-- internal
|   |-- handlers
|   |   |-- getWeather.go
|   |   |-- login.go
|   |-- services
|       |-- fetchWeather.go
|       |-- redisCacheData.go
|       |-- redisConnection.go
|-- pkg
|   |-- utils
|       |-- constants.go
|       |-- coordinatesPattern.go
|       |-- secConstants.go
|       |-- secHeader.go
|-- test
|   |-- cmd
|   |   |-- main_test.go
|   |-- internal
|   |   |-- handlers
|   |   |-- services
|   |-- pkg
|       |-- utils
|-- README.md
|-- go.mod
|-- go.sum
```

## Challenges in This Project

Eventhough I may not have a lot of experience with Go I try building an API as it was a preferred language for this Coding Challenge.

### Go Challenges

- Retaking the language.
- Researching methods to build an API
- Dependency Research
- Testing I have never made tests on Go.

### Redis Challenges

- Researching what is Redis used for
- Implementing Redis on a Go Project
  - Creating a connection
  - Storing Data
  - Calling Cache Data

## Project Approach

1. Retake Go and begin to make some code.
2. Research via Google and AI dependencies that are being needed to create and API and be secure.
3. Research what is Redis and how it can be implemented to Go.
4. Start the coding the challenge and creating a Docker container in local to test the API with Redis.
5. A Token Generation Endpoint was created in order to simulate security and not being able to consume the weather Endpoint that easy.
6. Add Security Headers to prevent different vulnerabilities such clickjacking or XSS attacks
7. Cleaning the code and moving strings to Constants.
8. Moving Subfunctions to utils as this can be re-used for other projects
9. Start creating unit testing.
10. Creating this documentation.

## Dependencies

- **Echo**: Is a high-performance, extensible, minimalist web framework for Go.
- **Echo JWT**: Middleware for Echo to handle JWT authentication.
- **Redis**: Used for caching weather data to improve performance and reduce the number of API calls to the NWS API.
- **Testify**: A toolkit with common assertions and mocks that plays nicely with the standard library.

## Endpoints

- **POST /login**: Generates a JWT token for authentication.
- **GET /weather**: Fetches weather data based on latitude and longitude. Requires a valid JWT token.

## Middleware

- **Security Headers**: Adds security-related headers to all responses.
- **Logger**: Logs HTTP requests.
- **Recover**: Recovers from panics and returns a 500 status code.

## Testing

The project includes unit tests for all major components. The tests are located in the `test` directory and cover the following:

- **Handlers**: Tests for the `getWeather` and `login` handlers to ensure they handle requests and responses correctly.
- **Services**: Tests for the `fetchWeather`, `redisCacheData`, and `redisConnection` services to ensure they interact with external APIs and Redis correctly.
- **Utils**: Tests for utility functions like `coordinatesPattern` and middleware like `securityHeaders`.

### What is on scope for the test?

- **Handlers**: Ensure that the handlers correctly parse request parameters, validate input, call the appropriate services, and return the correct HTTP status codes and responses.
- **Services**: Ensure that the services correctly interact with external APIs and Redis, handle errors gracefully, and return the expected data.
- **Utils**: Ensure that utility functions and middleware work as expected and add the necessary headers to responses.

## What Would be Next Steps?

For next steps is if this would be treated as a project that is going to production in this case a cloud project in GCP.

- Add SecConstans to a Secret Manager in order to protect the endpoints, signing key and ports.

- Create a DB where users could be stored and consumed by the login endpoint and review if the users exists and generating a new token.

- Create a FrondEnd where user can view the weather information.

- Create a DockerFile to upload the project to GCP and start working with the infrastructure.
