package urlify

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gregdel/pushover"
)

// Handle will handle the cloud function call
func Handle(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("URL")
	allowedRequester := os.Getenv("ALLOWED_REQUESTER")
	token := os.Getenv("TOKEN")
	user := os.Getenv("USER")

	resp, err := http.Get(url)

	if err != nil {
		log.Println("some error when fetching data")
		log.Println(r)
		return
	}

	defer resp.Body.Close()

	allowed := ""
	requester := r.Header.Get("X-Forwarded-For")
	if allowedRequester != requester {
		allowed = "not"
		log.Printf("Requester not allowed: %s", requester)
		fmt.Fprint(w, "Not Allowed")
	} else {
		log.Printf("Requester allowed: %s", requester)
		io.Copy(w, resp.Body)
	}

	app := pushover.New(token)
	recipient := pushover.NewRecipient(user)
	message := pushover.NewMessage(fmt.Sprintf("%s allowed request for decryption key from: %s", allowed, requester))
	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}
	log.Println(response)
}
