package remotemanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/michaeldcanady/omnimanageragent/internal/policy"
)

// RemoteManager is a struct that holds the configuration for making HTTP requests to the API endpoint.
type RemoteManager struct {
	URL    string
	Client *http.Client
}

// NewRemoteManager creates a new instance of RemoteManager with the provided URL and custom HTTP client settings.
func NewRemoteManager(url string) *RemoteManager {

	return &RemoteManager{
		URL: url,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetDeviceConfigurations fetches the device configurations from the API endpoint and returns them as a slice of Configuration.
func (m *RemoteManager) GetDeviceConfigurations() ([]policy.Configuration, error) {
	data, err := m.Get("/devicemangement/profiles") // Endpoint the fetch policies intended for a device
	if err != nil {
		return nil, fmt.Errorf("device configurations: %w", err)
	}

	var configurations []policy.Configuration
	err = json.Unmarshal(data, &configurations)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON data: %w", err)
	}

	return configurations, nil
}

// Get performs a GET request to the API endpoint and returns the response data.
func (m *RemoteManager) Get(endpoint string) ([]byte, error) {
	url := m.URL + endpoint
	resp, err := m.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET request returned non-OK status code: %d", resp.StatusCode)
	}

	body, err := readResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}

// readResponseBody reads the response body and returns it as a byte slice.
func readResponseBody(resp *http.Response) ([]byte, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
