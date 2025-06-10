package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {
    if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
	url = "http://" + url 
	fmt.Fprintf(os.Stderr, "Adding 'http://' to URL: %s\n", url)
	}
		resp, err := http.Get(url)
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
       fmt.Printf("Status code: %d\n", resp.StatusCode)
	   fmt.Printf("HTTp Status: %s\n", resp.Status)
		// The http.Get response body is an io.ReadCloser. We need to read it all.
		// b, err := io.ReadAll(resp.Body) // <--- CHANGE: Use io.ReadAll instead of ioutil.ReadAll
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close() // Don't forget to close the body to prevent resource leaks.
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		
		// fmt.Printf("loading the site %s", b)
	
}
}