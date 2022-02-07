package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var (
	createWatchlist          = "/user/v2/watchlist/create"
	renameWatchlist          = "/user/v2/watchlist/rename"
	deleteWatchlist          = "/user/v2/watchlist/delete"
	addEntryToWatchlist      = "/user/v2/watchlist/addEntry"
	deleteEntryFromWatchlist = "/user/v2/watchlist/removeEntry"
)

func (c *Client) CreateWatchlist(req models.CreateWatchlistRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+createWatchlist)
}

func (c *Client) RenameWatchlist(req models.RenameWatchlistRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+renameWatchlist)
}

func (c *Client) DeleteWatchlist(req models.DeleteWatchlistRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+deleteWatchlist)
}

func (c *Client) AddEntryToWatchlist(req models.EntryToWatchlistRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+addEntryToWatchlist)
}

func (c *Client) DeleteEntryFromWatchlist(req models.EntryToWatchlistRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+deleteEntryFromWatchlist)
}
