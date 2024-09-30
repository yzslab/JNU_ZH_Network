package utils

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"log"
	"net"
	"net/http"
	"time"
)

func createRestryClient(nic string) *resty.Client {
	client := resty.New()

	httpTransport := &http.Transport{
		TLSHandshakeTimeout:   3 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	if nic != "" {
		bind_addr := getNICAddress(nic)
		d := net.Dialer{
			LocalAddr: bind_addr,
			Timeout:   3 * time.Second,
		}
		httpTransport.Dial = d.Dial
	}

	client.SetTransport(httpTransport)

	return client
}

func getNICAddress(nic string) *net.TCPAddr {
	ief, err := net.InterfaceByName(nic)
	if err != nil {
		log.Fatal(err)
	}
	addrs, err := ief.Addrs()
	if err != nil {
		log.Fatal(err)
	}

	return &net.TCPAddr{
		IP: addrs[0].(*net.IPNet).IP,
	}
}
