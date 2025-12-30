package handlers

import (
	"fmt"
	"main/helpers"
	"main/models"
	"os"
	"path/filepath"
	"sync"

	"github.com/charmbracelet/log"
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
	// Log folder oluştur
	logPath := filepath.Join("log")
	err := os.MkdirAll(logPath, 0755)
	if err != nil {
		fmt.Println("Log folder couldn't be created:", err)
		return
	}

	// Log dosyası oluştur
	filePath := filepath.Join(logPath, "scan_report.log")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Report file couldn't be created:", err)
		return
	}
	defer file.Close()

	for _, res := range Results {
		line := fmt.Sprintf("[%s] Scanning: %s\n", res.Status, res.URL)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to report file:", err)
			return
		}
	}

	log.Info("Report saved:", filePath)
}
