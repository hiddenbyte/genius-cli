package useragent

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"jaytaylor.com/html2text"
)

func init() {
	userAgents = append(userAgents, geniusLyricsUserAgent{
		geniuslyricsAddr: regexp.MustCompile("https://genius.com/.+-lyrics")})
}

type geniusLyricsUserAgent struct {
	geniuslyricsAddr *regexp.Regexp
}

// Open renders the a genius lyrics HTML document
func (ua geniusLyricsUserAgent) Open(url string) error {
	// Request HTML document
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Render HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	lyricsHTML, err := doc.Find("div.lyrics").Html()
	if err != nil {
		return err
	}

	lyricsText, err := html2text.FromString(lyricsHTML)
	if err != nil {
		return err
	}

	fmt.Println(lyricsText)
	return nil
}

// Supports returns true, if the the address matches with https://genius.com/.+-lyrics
func (ua geniusLyricsUserAgent) Supports(address string) bool {
	return ua.geniuslyricsAddr.MatchString(address)
}
