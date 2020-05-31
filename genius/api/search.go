package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// SearchMetaResponse Genius API search meta response
type SearchMetaResponse struct {
	Response SearchResponse `json:"response"`
}

// SearchResponse Genius API search response
type SearchResponse struct {
	Hits []SearchHit `json:"hits"`
}

// SearchHit Genius API serch hit model
type SearchHit struct {
	Type   string                 `json:"type"`
	Result map[string]interface{} `json:"result"`
}

// Search searchs for documents hosted on Genius
func Search(q string, token string) (*SearchMetaResponse, error) {
	url := searchURL(q, token)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	searchResponse := &SearchMetaResponse{}
	err = json.NewDecoder(response.Body).Decode(searchResponse)
	if err != nil {
		return nil, err
	}
	return searchResponse, nil
}

// searchURL generates a Genius API search URL with the specifed search term and access token
func searchURL(q string, token string) string {
	return fmt.Sprintf("https://api.genius.com/search?access_token=%s&q=%s", token, url.QueryEscape(q))
}
