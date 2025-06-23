package util

import (
	"fmt"
	"net"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

var validate = validator.New()

type SystemInfo struct {
	Hostname     string `json:"hostname"`
	Domain       string `json:"domain,omitempty" envconfig:"DOMAIN"`
	FQDN         string `json:"fqdn,omitempty"`
	TestV6       string `json:"test_v6,omitempty"`
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

	addrs, err := net.LookupIP(si.TestEndpoint)
	if err != nil {
		// can't resolve FQDN, return err
		e := fmt.Errorf("failed to resolve host %s: %w", si.TestEndpoint,
			err)
		Logger.Err(e)
		return nil, e
	}
	// pull out a v6 addr from records
	for _, a := range addrs {
		if err := validate.Var(a.String(), "ipv6"); err != nil {
			si.TestV6 = a.String()
		}
	}

	Logger.Debug().Msgf("Node (FQDN): %s - Location: %s - Test Endpoint: %s - Test V6: %s", 
	si.FQDN, si.Location, si.TestEndpoint, si.TestV6)

	return &si, nil
}
