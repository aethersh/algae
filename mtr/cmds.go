package mtr

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/aethersh/algae/util"
)

func cmdPreamble(ip string, host string, cmd string) []byte {
	rtn := []byte("")
	if ip != host {
		rtn = strconv.AppendQuote(rtn, (fmt.Sprintln("Hostname %s resolved to %w", host, ip)))
	}
	rtn = strconv.AppendQuote(rtn, (fmt.Sprintln("$ %s", cmd)))
	return rtn
}

func RunPingCmd(host string) (*string, error) {
	ip, hostn, err := ValidateIPv6Host(host)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("ping", "-c 5", ip.String())
	cmdOut, err := cmd.Output()
	out := string(fmt.Appendln(cmdPreamble(ip.String(), *hostn, cmd.String()), cmdOut))

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

	cmd := exec.Command("mtr", "-6", "-z", "-b", "-r", "-w", "-o SLAVM", *hostn)
	cmdOut, err := cmd.Output()
	out := string(fmt.Appendln(cmdPreamble(ip.String(), *hostn, cmd.String()), cmdOut))

	if err != nil {
		util.Logger.Err(err).Msg("Failed to run MTR command")
		return &out, err
	}
	return &out, nil
}
