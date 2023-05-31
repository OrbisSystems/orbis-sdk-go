package model

import (
	"time"

	"github.com/google/uuid"
)

type NewsRequest struct {
	ID string `json:"id"`
}

type NewsResponse struct {
	News  []News `json:"news"`
	Count int64  `json:"count"`
}

type NewsData struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Title string `json:"title,omitempty"`
}

type News struct {
	ID             uuid.UUID        `json:"id"`
	SourceID       int64            `json:"source_id"`
	Author         string           `json:"author"`
	Created        time.Time        `json:"created"`
	Updated        time.Time        `json:"updated"`
	Title          string           `json:"title"`
	Teaser         string           `json:"teaser"`
	Body           string           `json:"body"`
	Url            string           `json:"url"`
	IsPressRelease bool             `json:"is_press_release"`
	ExtraData      map[int]NewsData `json:"data"`

	Image    []NewsImage   `json:"image"`
	Channels []NewsChannel `json:"channels"`
	Symbols  []NewsStock   `json:"symbols"`
	Tags     []NewsTag     `json:"tags"`
}

type NewsImage struct {
	Size string `json:"size"`
	Url  string `json:"url"`
}

type NewsChannel struct {
	Name string `json:"name"`
}

type NewsStock struct {
	Name string `json:"name"`
}

type NewsTag struct {
	Name string `json:"name"`
}

type HelpRequestWithSymbol struct {
	Symbol *string `json:"symbol"`
}

type NewsFilterRequest struct {
	Symbol       *string  `json:"symbol"`
	Author       *string  `json:"author"`
	StartDate    *string  `json:"start_date"`
	EndDate      *string  `json:"end_date"`
	Channels     []string `json:"channels"`
	Tags         []string `json:"tags"`
	PressRelease *bool    `json:"press_release"`
	Headline     *bool    `json:"headline"`

	Paging Paging `json:"paging"`
}
