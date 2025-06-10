package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		go fetch(url, ch) // Start a goroutine for each URL
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // Receive from the channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
func sanitizeFilename(url string) string {
	// Remove protocol and common URL components for a cleaner filename
	s := strings.TrimPrefix(url, "http://")
	s = strings.TrimPrefix(s, "https://")
	s = strings.TrimPrefix(s, "www.")

	// Replace non-alphanumeric, non-dot, non-dash characters with underscores
	reg := regexp.MustCompile(`[^a-zA-Z0-9\._-]+`)
	s = reg.ReplaceAllString(s, "_")

	// Trim any leading/trailing underscores
	s = strings.Trim(s, "_")

	// Ensure it has a default .html extension if none is clearly present
	if !strings.Contains(s, ".") {
		s += ".html"
	}

	// Limit length to prevent issues on some file systems
	if len(s) > 200 {
		s = s[:200]
	}

	return s
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	outputDirectory := "fetched_web_content"
	if err := os.MkdirAll(outputDirectory, 0755); err != nil {
		ch <- fmt.Sprintf("Error creating output directory: %v", err)
		resp.Body.Close()
		return
	}
	fileName := sanitizeFilename(url)
	filePath := filepath.Join(outputDirectory, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		ch <- fmt.Sprintf("Error creating file %s: %v", filePath, err)
		resp.Body.Close()
		return
	}
	defer file.Close()
	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // Close the response body to prevent resource leaks
	if err != nil {
		ch <- fmt.Sprintf("fetch: reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s (saved to %s)", secs, nbytes, url, filePath)

}
