package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var (
	journalRequest = "/transfer/journal/request"
	journalSearch  = "/transfer/journal/search"
)

func (c *Client) JournalRequest(req models.JournalRequest) (io.ReadCloser, error) {
	url := c.config.ApiUrl + journalRequest
	return c.client.DoPostAndReturnBody(req, url)
}

func (c *Client) JournalSearch(req models.JournalSearchRequest) (io.ReadCloser, error) {
	url := c.config.ApiUrl + journalSearch
	return c.client.DoPostAndReturnBody(req, url)
}
