package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/hiddenyte/genius-cli/genius/oauth"

	"github.com/hiddenyte/genius-cli/genius/api"
	"github.com/hiddenyte/genius-cli/genius/useragent"
)

// songFlagVar --song flag value
var songFlagVar string

// setupFlagVar --setup flag presence
var setupFlagVar bool

func init() {
	flag.StringVar(&songFlagVar, "song", "", "Prints the lyrics from a Genius song document")
	flag.StringVar(&songFlagVar, "s", "", "Prints the lyrics from a Genius song document (short)")
	flag.BoolVar(&setupFlagVar, "setup", false, "Setup genius-cli")
}

func main() {
	flag.Parse()

	if songFlagVar == "" && !setupFlagVar {
		flag.Usage()
		return
	}

	if setupFlagVar {
		setup()
	}

	if songFlagVar != "" {
		renderLyrics(songFlagVar)
	}
}

// renderLyrics renders the lyrics of the specified song
func renderLyrics(songSearchTerm string) {
	token := readToken()

	searchResponse, err := api.Search(songSearchTerm, token)
	if err != nil {
		log.Fatalf("renderLyrics: %v", err)
	}

	firstSearchHit := findFirstSongSearchHit(searchResponse.Response.Hits)
	if firstSearchHit == nil {
		log.Printf("renderLyrics: Unable to find the song: %s", songSearchTerm)
		return
	}

	lyricsURL := firstSearchHit.Result[api.DocumentURL].(string)
	ua := useragent.GetUserAgent(lyricsURL)
	if ua == nil {
		log.Printf("renderLyrics: Unable to render the document: %s", lyricsURL)
		return
	}

	ua.Open(lyricsURL)
}

// findFirstSongSearchHit retrieve the first search hit of 'song' type. returns nil, if there are none.
func findFirstSongSearchHit(hits []api.SearchHit) *api.SearchHit {
	for i, hit := range hits {
		if hit.Type == api.DocumentTypeSong {
			return &hits[i]
		}
	}
	return nil
}

// setup setups genius-cli
func setup() {
	token, err := oauth.Authorize()
	if err != nil {
		log.Fatalf("setup: %v", err)
	}
	persistToken(token)
}

// persistToken persists the access token
func persistToken(token string) {
	err := ioutil.WriteFile("./token", []byte(token), 0600)
	if err != nil {
		log.Fatalf("persistToken: %v", err)
	}
}

// readToken read access token
func readToken() string {
	token, err := ioutil.ReadFile("./token")
	if err != nil {
		log.Fatalf("readToken: %v", err)
	}
	return string(token)
}
