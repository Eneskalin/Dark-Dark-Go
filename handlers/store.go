package handlers

import (
	"fmt"
	"github.com/charmbracelet/log"
	"main/helpers"
	"main/models"
	"os"
	"path/filepath"
	"sync"
)

var (
	Results []models.ScanResult
	Mu      sync.Mutex
)

func SaveHtml(url string, body string) error {
	folderName := helpers.Normalized(url)

	docPath := filepath.Join("docs", folderName)
	err := os.MkdirAll(docPath, 0755)
	if err != nil {
		fmt.Println("Folder couldn't create: %v\n", err)
		return err
	}
	filePath := filepath.Join(docPath, "index.html")
	err = os.WriteFile(filePath, []byte(body), 0644)
	if err != nil {
		fmt.Println("File couldn't write: %v\n", err)
		return err
	}
	log.Info("Target content saved:", filePath)

	return nil

}

func SaveReport() {
	file, err := os.Create("../log/scan_report.log")
	if err != nil {
		fmt.Println("Rapor dosyası oluşturulamadı:", err)
		return
	}
	defer file.Close()
	for _, res := range Results {
		line := fmt.Sprintf("[%s] Scanning: %s\n", res.Status, res.URL)
		file.WriteString(line)
	}

}
