package helpers

import (
	"fmt"
	"net/http"
	"time"
	"github.com/charmbracelet/log"
	"golang.org/x/net/proxy"
)


func GetTorClient() (*http.Client, error) {
	proxyAddr := "127.0.0.1:9050"

	dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create proxy dialer : %w", err)
	}

	httpTransport := &http.Transport{
		Dial: dialer.Dial,
	}

	client := &http.Client{
		Transport: httpTransport,
		Timeout:   time.Second * 90, 
	}
	resp, err := client.Get("https://check.torproject.org")
	if err != nil {
		return nil, fmt.Errorf("Couldn't Tor Connection: %w", err)
	}
	defer resp.Body.Close()

	log.Info("Tor server successfuly.")
	return client, nil
}