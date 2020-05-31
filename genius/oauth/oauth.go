package oauth

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

// Authorize asks the user to authorize to access Genius on his behalf
func Authorize() (string, error) {
	conf := &oauth2.Config{
		ClientID:    "Y3ly8hs3HXL5pb9rNo94DYoKWnaPHt7uOf5qGeE-FSpZFqasF4ByksVZhR6ULyeV",
		RedirectURL: "http://localhost:9191/oauth/redirect",
		Endpoint:    oauth2.Endpoint{AuthURL: "https://api.genius.com/oauth/authorize"}}

	url := conf.AuthCodeURL("genius-cli", oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("response_type", "token"))

	// Start a HTTP server to handle the redirect uri
	http.HandleFunc("/oauth/redirect", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Content-Type", "text/plain")
		resp.WriteHeader(http.StatusOK)
		content := strings.NewReader("Copy the access token from the URL hash fragment and paste it into genius-cli.")
		content.WriteTo(resp)
	})
	go http.ListenAndServe(":9191", nil)
	fmt.Printf("Open the auth url %v\n", url)

	// Wait for the user to open the auth URL and paste the access token
	fmt.Print("Paste the 'access_token': ")
	reader := bufio.NewReader(os.Stdin)
	token, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return token[0 : len(token)-1], nil
}
