package net

import (
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func Zip(logger *zap.Logger, destination *os.File, url string) error {
	resp, err := Get(logger, url)
	if err != nil {
		logger.Error("Error making http call", zap.String("URL", url), zap.Error(err))
		return fmt.Errorf("Error making http call: %v", err)
	}
	defer resp.Body.Close()
	_, err = destination.ReadFrom(resp.Body)
	if err != nil {
		logger.Error("failed to write to file", zap.String("filePath", destination.Name()), zap.Error(err))
		return fmt.Errorf("failed to write to file %s: %v", destination.Name(), err)
	}
	return nil
}

func Get(logger *zap.Logger, url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("failed to download file", zap.String("URL", url), zap.Error(err))
		return nil, fmt.Errorf("failed to download file from %s: %v", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		logger.Error("Response received bad status", zap.String("StatusCode", resp.Status), zap.Error(err))
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}
	return resp, nil
}
