package orbis_api

type CreateWatchlistRequest struct {
	Title string `json:"title"`
}

type RenameWatchlistRequest struct {
	Title string `json:"title"`
	ID    int64  `json:"id"`
}

type DeleteWatchlistRequest struct {
	ID int64 `json:"id"`
}

type EntryToWatchlistRequest struct {
	WatchlistID int64  `json:"watchlistId"`
	Symbol      string `json:"symbol"`
	Option      bool   `json:"option"`
}
