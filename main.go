package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var version = "dev"

type App struct {
	Port     string
	Hostname string
	IP       string
	Version  string
}

func (a *App) Start() {
	addr := fmt.Sprintf(":%s", a.Port)
	log.Printf("Starting app on %s", addr)

	a.Hostname, a.IP = getHostInfo()

	http.Handle("/", logreq(a.index))

	log.Fatal(http.ListenAndServe(addr, nil))
}

func env(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}

func logreq(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path: %s", r.URL.Path)
		f(w, r)
	})
}

func (a *App) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hostname: %s\n", a.Hostname)
	fmt.Fprintf(w, "IP Address: %s\n", a.IP)
	fmt.Fprintf(w, "App Version: %s\n", version)
}

func getHostInfo() (string, string) {
	hostname, _ := os.Hostname()         // get hostname
	addrs, err := net.LookupIP(hostname) // returns a slice of the IP addresses of the host
	// lookupIP looks up host using the local resolver. It returns a slice of that host's IPv4 and IPv6 addresses.

	if err != nil {
		log.Println("Failed to detect machine host name. ", err.Error())
		return "unknown", "unknown"
	}

	var ip string
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			// only take the class B address
			if ipv4[0] == 172 {
				ip = ipv4.String()
				break
			}
		}
	}

	return hostname, ip
}

func main() {
	server := App{
		Port:    env("PORT", "8080"),
		Version: env("APP_VERSION", "1.0.0"),
	}
	server.Start()
}
