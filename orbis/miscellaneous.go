package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"time"
)

var (
	getMarketDate       = "/research/marketDate"
	getDatesMarket      = "/research/dates/markets"
	getLastOpen         = "/research/dates/lastOpen"
	getDatesCheck       = "/research/dates/check"
	getAdvisoryInfo     = "/v2/advisory/info"
	updateAdvisoryInfo  = "/v2/advisory/info/update"
	avatarUpload        = "/v1/files/user/avatar/upload"
	upload              = "/v1/files/upload"
	userAvatarUpload    = "/v1/files/upload/avatar/%s"
	getAvatar           = "/v1/files/user/avatar"
	getUserAvatar       = "/v1/files/%s/avatar"
	filesDownload       = "/v1/files/download/%s/%s"
	getFilesList        = "/v1/files/list"
	filesAllowedExts    = "/v1/files/allowed/exts"
	getErrorCodeListing = "/general/errors"
	ping                = "/general/ping"
)

func (c *Client) GetMarketDate() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getMarketDate, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetDatesMarket() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getDatesMarket, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetLastOpen(market string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getLastOpen, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("market", market)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetDatesCheck(market *string, date *time.Time) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getDatesCheck, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	if market != nil {
		q.Add("market", *market)
	}
	if date != nil {
		q.Add("date", date.Format("01/02/2006"))
	}
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAdvisoryInfo() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getAdvisoryInfo, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) UpdateAdvisoryInfo(req models.AdvisoryInfoUpdateRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+updateAdvisoryInfo)
}

func (c *Client) AvatarUpload(req models.AvatarUploadRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+updateAdvisoryInfo)
}

func (c *Client) Upload(req models.UploadRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+updateAdvisoryInfo)
}

func (c *Client) UserAvatarUpload(ID string, req models.AvatarUploadRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+fmt.Sprintf(userAvatarUpload, ID))
}

func (c *Client) GetAvatar() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getAvatar, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetUserAvatar(ID string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getUserAvatar, ID), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) FilesDownload(tag, ID string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(filesDownload, tag, ID), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetFilesList(tag, fileName string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getFilesList, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("tag", tag)
	q.Add("fileName", fileName)

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) FilesAllowedExts() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+filesAllowedExts, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetErrorCodeListing() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getErrorCodeListing, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) Ping() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+ping, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
