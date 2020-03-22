package pfsense

import (
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"strings"
)

func (pfsense *Pfsense) getToken(url string) string {
	resp, err := pfsense.Client.Get(url)
	if err != nil {
		log.Fatalf("Failed get token: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed get token: %v", err)
	}

	regex := regexp.MustCompile(`var csrfMagicToken = "([^"]+)";`)
	match := regex.FindStringSubmatch(string(body))
	if len(match) < 1 {
		log.Fatalf("Failed find csrf token: %v", err)
	}
	token := match[1]
	return token
}

func (pfsense *Pfsense) Login() {
	currentUrl := pfsense.Settings.Url
	token := pfsense.getToken(currentUrl)
	data := url.Values{
		"__csrf_magic": {token},
		"usernamefld":  {pfsense.Settings.User},
		"passwordfld":  {pfsense.Settings.Pass},
		"login":        {"Sign In"},
	}
	_, err := pfsense.Client.Post(
		currentUrl,
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}
}
