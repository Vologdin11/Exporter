package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"stp-exporter/internal/config"
)

type token struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
	Token_type   string `json:"token_type"`
	Scope        string `json:"scope"`
}

type metrics struct {
	Columns []string        `json:"columns"`
	Values  [][]interface{} `json:"values"`
}

func GetAllMetrics(config config.Config, token token) ([]metrics, error) {
	metrics := make([]metrics, len(config.Tables))
	for i, table := range config.Tables {
		metric, err := GetMetric(table.Name, token)
		if err != nil {
			return nil, err
		}
		metrics[i] = metric
	}
	return metrics, nil
}

func GetMetric(table string, token token) (metrics, error) {
	request, err := createMetricRequest(table, token)
	if err != nil {
		return metrics{}, err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return metrics{}, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return metrics{}, err
	}
	newMetrics := metrics{}
	err = json.Unmarshal(body, &newMetrics)
	if err != nil {
		return metrics{}, err
	}
	return newMetrics, nil
}

func createMetricRequest(table string, token token) (*http.Request, error) {
	env, err := getEnv("DB", "URL")
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/viqube/databases/%s/tables/%s/records", env["URL"], env["DB"], table), nil)
	if err != nil {
		return &http.Request{}, err
	}
	request.Header.Set("X-API-VERSION", "3.5")
	request.Header.Set("Authorization", token.Token_type+token.Access_token)
	return request, nil
}

func GetToken() (token, error) {
	client := http.Client{}
	request, err := createTokenRequest()
	if err != nil {
		return token{}, err
	}
	response, err := client.Do(request)
	if err != nil {
		return token{}, err
	}
	defer response.Body.Close()
	bodyToken, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return token{}, err
	}
	newToken := token{}
	err = json.Unmarshal(bodyToken, &newToken)
	if err != nil {
		return token{}, err
	}
	return newToken, nil
}

func createTokenRequest() (*http.Request, error) {
	env, err := getEnv("URL", "LOGIN", "PASSWORD", "AUTHORIZATION")
	if err != nil {
		return nil, err
	}
	contentType := "application/x-www-form-urlencoded"
	bodyDataRaw := fmt.Sprintf("grant_type=password&scope=viqube_api+viqubeadmin_api&response_type=id_token+token&username=%s&password=%s", env["LOGIN"], env["PASSWORD"])

	requestBody, err := json.Marshal(map[string]string{
		"data-raw": bodyDataRaw,
	})
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", env["URL"], bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Authorization", env["AUTHORIZATION"])
	return request, nil
}

func getEnv(envNames ...string) (map[string]string, error) {
	env := make(map[string]string, len(envNames))
	for _, envName := range envNames {
		tmp := os.Getenv(envName)
		if tmp == "" {
			return nil, fmt.Errorf("variable %s not set", envName)
		}
		env[envName] = tmp
	}
	return env, nil
}
