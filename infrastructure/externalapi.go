package infrastructure

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/hikata101/health_data_service/logger"
)

var who_url string = "https://dw.euro.who.int/api/v3/"

func Execute(ctx context.Context, indicatorCode string, query string) (string, error) {
	logger.Logger.Debug("Executing WHO Europe API request")
	resp := ""
	client := &http.Client{}
	url := who_url + "/export/" + indicatorCode + "?format=csv&" + query
	reqHTTP, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		logger.Logger.Error("failed to create request: " + err.Error())
		return "", err
	}
	var data struct {
		StatusURL   string `json:"check_status_url"`
		ExportURL   string `json:"export_url"`
		Status      string `json:"status"`
		DownloadURL string `json:"download_url"`
	}
	res, err := client.Do(reqHTTP)
	if err != nil {
		logger.Logger.Error("failed to execute request: " + err.Error())
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Logger.Error("failed to read response: " + err.Error())
		return "", err
	}

	resp = string(body)
	logger.Logger.Debug("initial request response: " + resp)
	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		logger.Logger.Error("failed to parse response JSON: " + err.Error())
		return "", err
	}
	for {
		if data.StatusURL == "" {
			break
		}
		reqHTTP, err := http.NewRequestWithContext(ctx, "GET", data.StatusURL, nil)
		if err != nil {
			logger.Logger.Error("failed to create request: " + err.Error())
			return "", err
		}
		res, err := client.Do(reqHTTP)
		if err != nil {
			logger.Logger.Error("failed to execute request: " + err.Error())
			return "", err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			logger.Logger.Error("failed to read response: " + err.Error())
			return "", err
		}

		resp = string(body)
		logger.Logger.Debug("initial request response: " + resp)
		if err := json.Unmarshal([]byte(resp), &data); err != nil {
			logger.Logger.Error("failed to parse response JSON: " + err.Error())
			return "", err
		}
		if data.Status == "InQueue" || data.Status == "Processing" {
			time.Sleep(2 * time.Second)
			continue
		}
		if data.Status == "Failed" {
			logger.Logger.Error("data generation failed")
			return "", err
		}
		if data.Status == "Generated" {
			break
		}
		logger.Logger.Error("unknown status: " + data.Status)
		return "", err
	}

	logger.Logger.Debug("data generated, proceeding to download")
	downloadURL := data.DownloadURL
	reqHTTP, err = http.NewRequestWithContext(ctx, "GET", downloadURL, nil)
	if err != nil {
		logger.Logger.Error("failed to create request: " + err.Error())
		return "", err
	}

	res, err = client.Do(reqHTTP)
	if err != nil {
		logger.Logger.Error("failed to execute request: " + err.Error())
		return "", err
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		logger.Logger.Error("failed to read response: " + err.Error())
		return "", err
	}
	resp = string(body)
	logger.Logger.Debug("response status: " + res.Status)
	logger.Logger.Debug("WHO Europe API request executed successfully")
	return resp, nil
}

func ExecuteTest() (string, error) {
	logger.Logger.Debug("ExecutingTest WHO Europe API request")
	client := &http.Client{}
	url := who_url + "/countries"
	reqHTTP, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Logger.Error("failed to create request: " + err.Error())
		return "", err
	}
	res, err := client.Do(reqHTTP)
	if err != nil {
		logger.Logger.Error("failed to execute request: " + err.Error())
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Logger.Error("failed to read response: " + err.Error())
		return "", err
	}

	resp := string(body)
	logger.Logger.Debug("response status: " + res.Status)
	logger.Logger.Debug("WHO Europe API request executed successfully")
	return resp, nil
}
