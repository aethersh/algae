package mtr

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/aethersh/algae/util"
	"github.com/gofiber/fiber/v2"
)

func cmdPreamble(ip string, host string, cmd string) string {
	rtn := ""
	if ip != host {
		rtn = fmt.Sprintf("Hostname %s resolved to %s\n", host, ip)
	}
	cmdParts := strings.Split(cmd, "/")
	if strings.Contains(ip, "/") {
		cmd = cmdParts[len(cmdParts)-2]
	} else {
		cmd = cmdParts[len(cmdParts)-1]
	}

	rtn = fmt.Sprintf("%s$ %s", rtn, cmd)
	return rtn
}

func RunPingCmd(host string) (*string, error) {
	ip, hostn, err := ValidateIPv6Host(host)
	if err != nil {
		return nil, err
	}

	util.Logger.Debug().Msgf("Pinging address %s", ip)

	cmd := exec.Command("ping6", "-c 5", ip.String())
	cmdOut, err := cmd.Output()
	out := fmt.Sprintf("%s\n%s", cmdPreamble(ip.String(), *hostn, cmd.String()), string(cmdOut))

	if err != nil {
		util.Logger.Err(err).Msg("Failed to run ping command")
		return &out, err
	}
	return &out, nil
}

func RunMTRCmd(host string) (*string, error) {
	ip, hostn, err := ValidateIPv6Host(host)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("mtr", "-6", "-z", "-b", "-r", "-w", "-o SLAVM", ip.String())
	cmdOut, err := cmd.Output()
	out := fmt.Sprintf("%s\n%s", cmdPreamble(ip.String(), *hostn, cmd.String()), string(cmdOut))

	if err != nil {
		util.Logger.Err(err).Msg("Failed to run MTR command")
	}
	return &out, nil
}

func RunBIRDCmd(cidr string) (*string, int, error) {
	valid := false

	if _, err := ValidateIPv6CIDR(cidr); err == nil {
		valid = true
	} else if _, err := ValidateIPv6Address(cidr); err == nil {
		valid = true
	}

	if !valid {
		res := "invalid IPv6 address or CIDR"
		return &res, fiber.StatusBadRequest, fmt.Errorf(res)
	}

	cmd := exec.Command("birdc", "show route all for", cidr)
	cmdOut, err := cmd.Output()
	out := fmt.Sprintf("%s\n%s", cmdPreamble(cidr, cidr, cmd.String()), string(cmdOut))
	if err != nil {
		util.Logger.Err(err).Msg("Failed to run birdc command")
		return &out, fiber.StatusBadRequest, err
	}

	return &out, fiber.StatusOK, nil
}
