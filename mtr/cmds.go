package mtr

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/aethersh/algae/util"
)

func cmdPreamble(ip string, host string, cmd string) string {
	rtn := ""
	if ip != host {
		rtn = fmt.Sprintf("Hostname %s resolved to %s\n", host, ip)
	}
	cmdParts := strings.Split(cmd, "/")
	cmd = cmdParts[len(cmdParts)-1]

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
