package util

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type SystemInfo struct {
	Hostname string `json:"hostname"`
	Domain   string `json:"domain,omitempty" envconfig:"DOMAIN"`
	FQDN     string `json:"fqdn,omitempty"`
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

	return &si, nil
}
