package util

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

var validate = validator.New()

type SystemInfo struct {
	Hostname     string `json:"hostname"`
	Domain       string `json:"domain,omitempty" envconfig:"DOMAIN"`
	FQDN         string `json:"fqdn,omitempty"`
	TestV6       string `json:"test_v6,omitempty" envconfig:"TEST_V6"`
	TestEndpoint string `json:"test_endpoint,omitempty"`
	Location     string `json:"location,omitempty" envconfig:"LOCATION"`
}

func GetSystemInfo() (*SystemInfo, error) {
	var si SystemInfo

	if err := envconfig.Process(ENVCONFIG_PREFIX, &si); err != nil {
		return nil, err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	si.Hostname = hostname
	si.FQDN = fmt.Sprintf("%s.%s", si.Hostname, si.Domain)
	si.TestEndpoint = fmt.Sprintf("6.%s", si.FQDN)

	Logger.Debug().Msgf("Node (FQDN): %s - Location: %s - Test Endpoint: %s - Test V6: %s", 
	si.FQDN, si.Location, si.TestEndpoint, si.TestV6)

	return &si, nil
}
