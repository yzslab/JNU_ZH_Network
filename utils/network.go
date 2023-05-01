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

	if nic != "" {
		bind_addr := getNICAddress(nic)
		d := net.Dialer{
			LocalAddr: bind_addr,
			Timeout:   3 * time.Second,
		}
		client.SetTransport(&http.Transport{
			Dial:                d.Dial,
			TLSHandshakeTimeout: 3 * time.Second,
		})
	}

	// must place after client.SetTransport()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

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
