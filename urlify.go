package urlify

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Handle will handle the cloud function call
func Handle(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("URL")
	resp, err := http.Get(url)

	if err != nil {
		log.Println("some error when fetching data")
		return
	}

	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}
