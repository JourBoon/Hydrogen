package pkg

import (
	"fmt"
	"os"
	"encoding/json"
)

func LoadConfig(chemin string) (*ProxyConfig, error) {
	file, err := os.Open(chemin)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config ProxyConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func GetConfigValue(config *ProxyConfig, key string) (interface{}, error) {
	switch key {
	case "proxy.listen_address":
		return config.Proxy.ListenAddress, nil
	case "proxy.listen_port":
		return config.Proxy.ListenPort, nil
	case "proxy.backend_address":
		return config.Proxy.BackendAddress, nil
	case "proxy.backend_port":
		return config.Proxy.BackendPort, nil
	case "security.rate_limit":
		return config.Security.RateLimit, nil
	case "security.timeout_seconds":
		return config.Security.TimeoutSeconds, nil
	default:
		return nil, fmt.Errorf("Clé non trouvée: %s", key)
	}
}