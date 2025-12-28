package controllers

import (
	"net/http"
	"strings"
	"net"
)


func GetClientIP(r *http.Request) string {
	// Check standard headers for the real IP when behind a proxy or load balancer
	for _, headerName := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		ip := r.Header.Get(headerName)
		if ip != "" {
			// X-Forwarded-For can contain a list of IPs, the first one is the client's
			ips := strings.Split(ip, ",")
			for i, p := range ips {
				ips[i] = strings.TrimSpace(p)
			}
			return ips[0]
		}
	}

	// If no proxy headers are found, use the RemoteAddr field
	// RemoteAddr includes the port, so we use strings.Split to get only the IP part
	// We handle potential errors or alternative formats with net.SplitHostPort if needed
	if r.RemoteAddr != "" {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil {
			return ip
		}
		// Fallback for cases where SplitHostPort might fail (though rare for valid RemoteAddr values)
		return r.RemoteAddr
	}

	return ""
}
