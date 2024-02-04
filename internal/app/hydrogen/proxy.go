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
    proxy  *httputil.ReverseProxy
}

func NewProxy(config pkg.ProxyConfig) pkg.ReverseProxy {
    targetURL := fmt.Sprintf("http://%s:%d", config.Proxy.BackendAddress, config.Proxy.BackendPort)
    return &proxyImpl{
        config: config,
        proxy:  httputil.NewSingleHostReverseProxy(MustParseURL(targetURL)),
    }
}

func (pi *proxyImpl) HandleRequest(w http.ResponseWriter, r *http.Request) {
    pkg.Info("Request Specification - below:", r.Method, r.URL.Path)
    pi.proxy.ServeHTTP(w, r)
}

func MustParseURL(rawURL string) *url.URL {
    u, err := url.Parse(rawURL)
    if err != nil {
        panic(err)
    }
    return u
}