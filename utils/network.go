package utils

import (
	"net"
	"os/exec"
	"runtime"
	"strings"
)

// PingIP checks if a host is reachable by sending a single ping request.
func PingIP(host string) bool {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "-n", "1", "-w", "1000", host)
	} else {
		cmd = exec.Command("ping", "-c", "1", "-W", "1", host)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return strings.Contains(string(output), "TTL") || strings.Contains(string(output), "ttl")
}

// IsResolvable checks if a hostname can be resolved to an IP address.
func IsResolvable(host string) bool {
	ips, err := net.LookupIP(host)
	return err == nil && len(ips) > 0
}
