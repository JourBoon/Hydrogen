package pkg

import (
	"net/http"
)

type ReverseProxy interface {
    HandleRequest(w http.ResponseWriter, r *http.Request)
}


type ProxyConfig struct {
	Proxy struct {
	  ListenAddress string `json:"listen_address"`
	  ListenPort    int    `json:"listen_port"`
	  Backends      []struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	  } `json:"backends"`
	} `json:"proxy"`
	Security struct {
	  RateLimit      int `json:"rate_limit"`
	  TimeoutSeconds int `json:"timeout_seconds"`
	} `json:"security"`
  }
  