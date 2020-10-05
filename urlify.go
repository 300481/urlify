package urlify

import (
	"io"
	"log"
	"net/http"
)

// Handle will handle the cloud function call
func Handle(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://media.os.fressnapf.com/cms/2020/08/5f2cfc5f7f982-5f2cfc5f7f983Ratgeber-Katze_Outdoor_Kitten_1200x527.jpg.jpg?t=cmsimg_920")

	if err != nil {
		log.Println("some error when fetching data")
		return
	}

	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}
