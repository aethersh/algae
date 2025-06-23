package mtr

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateIPv6Address validates an IPv6 address and returns a net.IPAddr.
// It returns an error if the address is not valid.
func ValidateIPv6Address(ip string) (*net.IPAddr, error) {
	if err := validate.Var(ip, "ipv6"); err != nil {
		return nil, err
	}
	addr := net.ParseIP(ip)
	if addr == nil {
		return nil, fmt.Errorf("invalid IPv6 address: %s", ip)
	}
	if addr.IsUnspecified() || addr.IsLoopback() || addr.IsMulticast() {
		return nil, fmt.Errorf("address %s is not a valid IPv6 address", ip)
	}
	return &net.IPAddr{IP: addr}, nil
}

// ValidateIPv6Host takes an IPv6 address or FQDN. If it's an FQDN, it checks if it resolves to an IPv6 address.
// It returns a net.IPAddr if the host is valid, or an error if it is not.
func ValidateIPv6Host(host string) (ip *net.IPAddr, hostn *string, err error) {
	if ip, err := ValidateIPv6Address(host); err == nil {
		// Is valid IPv6 address, return success
		return ip, &host, nil
	}
	if err := validate.Var(host, "fqdn"); err != nil {
		// Not valid FQDN, return err
		return nil, nil, fmt.Errorf("invalid host: %s, error: %w", host, err)
	}
	// Is valid FQDN, lookup for v6 address
	addrs, err := net.LookupIP(host)
	if err != nil {
		// can't resolve FQDN, return err
		return nil, nil, fmt.Errorf("failed to resolve host %s: %w", host,
			err)
	}
	// pull out a v6 addr from records
	for _, a := range addrs {
		if ip, err := ValidateIPv6Address(a.String()); err == nil {
			return ip, &host, nil
		}
	}
	return nil, nil, fmt.Errorf("no valid IPv6 address found for host %s", host)
}

// ValidateIPv6CIDR validates an IPv6 CIDR notation and returns a net.IPNet.
func ValidateIPv6CIDR(cidr string) (*net.IPNet, error) {
	if err := validate.Var(cidr, "cidrv6"); err != nil {
		return nil, err
	}
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR notation: %s, error: %w", cidr, err)
	}
	// Check if the IPv6 CIDR Mask is a /48 or bigger
	maskInt, _ := strconv.Atoi((strings.Split(cidr, "/")[1]))
	if maskInt > 48 {
		return nil, fmt.Errorf("CIDR %s mask is too small (must be /48 or larger)", cidr)
	}
	return ipNet, nil
}
