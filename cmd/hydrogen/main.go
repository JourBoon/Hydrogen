// main.go
package main

import (
    "fmt"
	"net/http"
	"hydrogen/internal/app/hydrogen"
	"hydrogen/pkg"
)

func main() {
	config, err := pkg.LoadConfig("configs/config.json")
	if err != nil {
		fmt.Println("Error while reading configuration file:", err)
		return
	}

	listenAddress, err := pkg.GetConfigValue(config, "proxy.listen_address")
	if err != nil {
		fmt.Println("Error while fetching value:", err)
		return
	}

	listenPort, err := pkg.GetConfigValue(config, "proxy.listen_port")
	if err != nil {
		fmt.Println("Error while fetching value:", err)
		return
	}

	addr := fmt.Sprintf("%s:%d", listenAddress, listenPort)

	proxy := internal.NewProxy(*config)

	http.HandleFunc("/", proxy.HandleRequest)

	pkg.Info("Hydrogen listening on %s... \n Waiting for entries...", addr)
	http.ListenAndServe(addr, nil)
}
