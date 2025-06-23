package mtr_test

import (
	// "net"
	"testing"

	"github.com/aethersh/algae/mtr"
	"github.com/stretchr/testify/assert"
)

func TestValidateIPv6Address(t *testing.T) {
	tests := []struct {
		ip    string
		valid bool
	}{
		{"2602:fbcf:df::1", true},
		{"::", false},
		{"invalid-ip", false},
	}

	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			addr, err := mtr.ValidateIPv6Address(tt.ip)
			if tt.valid {
				assert.NoError(t, err)
				assert.NotNil(t, addr)
			} else {
				assert.Error(t, err)
				assert.Nil(t, addr)
			}
		})
	}
}

func TestValidateIPv6Host(t *testing.T) {
	tests := []struct {
		host  string
		valid bool
	}{
		{"henrikvt.com", true},
		{"anycast.as215207.net", true},
		{"github.com", false},
		{"2602:fbcf:df::1", true},
		{"::1", false},
	}

	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			addr, err := mtr.ValidateIPv6Host(tt.host)
			if tt.valid {
				assert.NoError(t, err)
				assert.NotNil(t, addr)
			} else {
				assert.Error(t, err)
				assert.Nil(t, addr)
			}
		})
	}
}

func TestValidateIPv6CIDR(t *testing.T) {
	tests := []struct {
		cidr  string
		valid bool
	}{
		{"henrikvt.com", false},
		{"2602:fbcf:d0::/47", true},
		{"2602:fbcf:df::/48", true},
		{"2602:fbcf:df::1/128", false},
		{"::1", false},
	}

	for _, tt := range tests {
		t.Run(tt.cidr, func(t *testing.T) {
			addr, err := mtr.ValidateIPv6CIDR(tt.cidr)
			if tt.valid {
				assert.NoError(t, err)
				assert.NotNil(t, addr)
			} else {
				assert.Error(t, err)
				assert.Nil(t, addr)
			}
		})
	}
}
