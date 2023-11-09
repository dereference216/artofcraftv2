package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func Write(filename string, content string, threadlock *sync.Mutex) {
	threadlock.Lock()
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Close()
		threadlock.Unlock()
	}()
	if _, err := io.WriteString(file, content+"\n"); err != nil {
		log.Printf("Error writing to file: %v\n", err)
	}
}

func ReadFile(filePath string, threadlock *sync.Mutex) ([]string, error) {
	threadlock.Lock()
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		file.Close()
		threadlock.Unlock()
	}()
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
