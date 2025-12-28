package helpers

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func ReadTargets(filePath string) ([]string, error) {
    var targets []string

    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("dosya açılırken hata oluştu: %w", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        cleanURL := strings.TrimSpace(line)
        
        if cleanURL != "" {
            targets = append(targets, cleanURL)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return targets, nil
}