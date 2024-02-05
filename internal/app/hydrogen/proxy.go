package internal

import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "net/url"
	"hydrogen/pkg"
)

type proxyImpl struct {
    config pkg.ProxyConfig
    proxies []*httputil.ReverseProxy
}

func NewProxy(config pkg.ProxyConfig) pkg.ReverseProxy {
    proxies := make([]*httputil.ReverseProxy, len(config.Proxy.Backends))
	for i, backend := range config.Proxy.Backends {
		targetURL := fmt.Sprintf("http://%s:%d", backend.Address, backend.Port)
		proxies[i] = httputil.NewSingleHostReverseProxy(pkg.MustParseURL(targetURL))
	}

	return &proxyImpl{
		config: config,
		proxies: proxies,
	}
}

func (pi *proxyImpl) HandleRequest(w http.ResponseWriter, r *http.Request) {
    pkg.Info("Request Specification - below:", r.Method, r.URL.Path)
    selectedProxy := pi.proxies[0]
    selectedProxy.ServeHTTP(w, r)
}

func MustParseURL(rawURL string) *url.URL {
    u, err := url.Parse(rawURL)
    if err != nil {
        panic(err)
    }
    return u
}