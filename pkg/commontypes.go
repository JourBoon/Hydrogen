package pkg

import (
	"net/http"
)

type ReverseProxy interface {
    HandleRequest(w http.ResponseWriter, r *http.Request)
}


type ProxyConfig struct {
	Proxy struct {
		ListenAddress  string `json:"listen_address"`
		ListenPort     int    `json:"listen_port"`
		BackendAddress string `json:"backend_address"`
		BackendPort    int    `json:"backend_port"`
	} `json:"proxy"`
	Security struct {
		RateLimit       int `json:"rate_limit"`
		TimeoutSeconds  int `json:"timeout_seconds"`
	} `json:"security"`
}