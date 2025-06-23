package mtr

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/aethersh/algae/util"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateIPv6Address validates an IPv6 address and returns a net.IPAddr.
// It returns an error if the address is not valid.
func ValidateIPv6Address(ip string) (*net.IPAddr, error) {
	ip = strings.TrimSpace(ip)
	if err := validate.Var(ip, "ipv6"); err != nil {
		util.Logger.Err(err).Msgf("not valid ipv6")
		return nil, err
	}
	addr := net.ParseIP(ip)
	if addr == nil {
		e := fmt.Errorf("invalid IPv6 address: %s", ip)
		util.Logger.Err(e).Msgf("could not parse ipv6")
		return nil, e
	}
	if addr.IsUnspecified() || addr.IsLoopback() || addr.IsMulticast() {
		e := fmt.Errorf("address %s is not a valid IPv6 address", ip)
		util.Logger.Err(e).Msgf("invalid ipv6 class")
		return nil, e
	}
	return &net.IPAddr{IP: addr}, nil
}

// ValidateIPv6Host takes an IPv6 address or FQDN. If it's an FQDN, it checks if it resolves to an IPv6 address.
// It returns a net.IPAddr if the host is valid, or an error if it is not.
func ValidateIPv6Host(host string) (ip *net.IPAddr, hostn *string, err error) {
	host = strings.TrimSpace(host)
	if ip, err := ValidateIPv6Address(host); err == nil {
		util.Logger.Debug().Msgf("Found valid host v6 address: %s", ip)
		// Is valid IPv6 address, return success
		return ip, &host, nil
	}

	if err := validate.Var(host, "fqdn"); err != nil {
		util.Logger.Err(err).Msgf("FQDN validation error")
		// Not valid FQDN, return err
		e := fmt.Errorf("invalid host: %s, error: %w", host, err)
		util.Logger.Err(e)
		return nil, nil, e
	}
	// Is valid FQDN, lookup for v6 address
	addrs, err := net.LookupIP(host)
	if err != nil {
		// can't resolve FQDN, return err
		e := fmt.Errorf("failed to resolve host %s: %w", host,
			err)
		util.Logger.Err(e)
		return nil, nil, e
	}
	// pull out a v6 addr from records
	for _, a := range addrs {
		if ip, err := ValidateIPv6Address(a.String()); err == nil {
			return ip, &host, nil
		}
	}

	e := fmt.Errorf("no valid IPv6 address found for host %s", host)
	util.Logger.Err(e)
	return nil, nil, e
}

// ValidateIPv6CIDR validates an IPv6 CIDR notation and returns a net.IPNet.
func ValidateIPv6CIDR(cidr string) (*net.IPNet, error) {
	if err := validate.Var(cidr, "cidrv6"); err != nil {
		util.Logger.Err(err)
		return nil, err
	}
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		e := fmt.Errorf("invalid CIDR notation: %s, error: %w", cidr, err)
		util.Logger.Err(e)
		return nil, e
	}
	// Check if the IPv6 CIDR Mask is a /48 or bigger
	maskInt, _ := strconv.Atoi((strings.Split(cidr, "/")[1]))
	if maskInt > 48 {
		e := fmt.Errorf("CIDR %s mask is too small (must be /48 or larger)", cidr)
		util.Logger.Err(e)
		return nil, e
	}
	return ipNet, nil
}
