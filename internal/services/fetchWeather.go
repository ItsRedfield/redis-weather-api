package services

import (
	"cloudflare-challenge-weaher-api/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchWeatherData(lat, lon string) (map[string]interface{}, error) {

	gridpointURL := fmt.Sprintf(utils.NWSAPI, lat, lon)
	resp, err := http.Get(gridpointURL)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFetchGPD, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(utils.ErrorInvalidCoord)
	}

	var gridpointData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&gridpointData); err != nil {
		return nil, fmt.Errorf(utils.ErrorDecodeGridP, err)
	}

	forecastURL, ok := gridpointData[utils.Properties].(map[string]interface{})["forecast"].(string)
	if !ok {
		return nil, fmt.Errorf(utils.ErrorURLNotFound)
	}

	resp, err = http.Get(forecastURL)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFetchFD, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(utils.ErrorFetchInvalidCoordOrLoc)
	}

	var forecastData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&forecastData); err != nil {
		return nil, fmt.Errorf(utils.ErrorDecodeFD, err)
	}

	properties, ok := forecastData[utils.Properties].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf(utils.ErrorPropertiesNotFoundFD)
	}

	periods, ok := properties["periods"].([]interface{})
	if !ok || len(periods) == 0 {
		return nil, fmt.Errorf(utils.ErrorNoPeriodsFound)
	}

	return periods[0].(map[string]interface{}), nil
}
