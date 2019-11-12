package pfsense

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

type Config struct {
	Content  []byte
	Filename string
}

func getFilename(response *http.Response) string {
	value := response.Header.Get("Content-Disposition")
	filename := strings.TrimPrefix(value, "attachment; filename=")
	return filename
}

func GetConfig(pfsense *Pfsense) *Config {
	currentUrl := pfsense.Settings.Url + "/diag_backup.php"
	token := getToken(pfsense, currentUrl)
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	defer w.Close()
	err := w.WriteField("__csrf_magic", token)
	if err != nil {
		log.Fatalf("Failed to create form field: %v", err)
	}
	err = w.WriteField("donotbackuprrd", "yes")
	if err != nil {
		log.Fatalf("Failed to create form field: %v", err)
	}
	err = w.WriteField("download", "Download configuration as XML")
	if err != nil {
		log.Fatalf("Failed to create form field: %v", err)
	}
	resp, err := pfsense.Client.Post(
		currentUrl,
		w.FormDataContentType(),
		buf,
	)
	if err != nil {
		log.Fatalf("Failed get config: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed get config: %v", err)
	}
	filename := getFilename(resp)
	config := &Config{
		Content:  body,
		Filename: filename,
	}
	return config
}
