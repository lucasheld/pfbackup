package pfsense

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
)

type Settings struct {
	Url      string
	User     string
	Pass     string
	NoVerify bool
}

type Pfsense struct {
	Settings *Settings
	Client   *http.Client
}

func InitClient(config *Settings) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: config.NoVerify},
	}
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Transport: tr,
		Jar:       cookieJar,
	}
	return client
}
