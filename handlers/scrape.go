package handlers

import (
	"io"
	"net/http"

	"github.com/charmbracelet/log"
)

func Screaper(client *http.Client, target string) error {

	req, err := http.NewRequest("GET", target, nil)
	if err != nil {

		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:109.0) Gecko/20100101 Firefox/115.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Error("Scanning failed", "url", target, "hata", err)
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Body okunamadı", "url", target, "hata", err)
		return err
	}
	err = SaveHtml(target, string(bodyBytes))
	if err != nil {

	}

	return nil
}
