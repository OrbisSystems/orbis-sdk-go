package model

import (
	"time"

	"github.com/google/uuid"
)

type NewsFilterRequest struct {
	Symbol         *string  `json:"symbol"`          // Symbol. See available values by using endpoint /symbols
	NewsSource     []string `json:"news_source"`     // News source. See available values by using endpoint /sources
	StartDate      *string  `json:"start_date"`      // Start date for searching news
	EndDate        *string  `json:"end_date"`        // End date for searching news
	Language       *string  `json:"language"`        // Which news language. See available values by using endpoint /lang
	RelevanceLevel *string  `json:"relevance_level"` // Relevance level. Required with subject				Enums(week, moderate, relevant, strongly)
	Subject        []string `json:"subject"`         // News subjects. Required with relevance_level 		Enums(IS/biz, IS/culture, IS/fin etc. See available values by using endpoint /taxonomy)
	Paging         *Paging  `json:"paging"`          // Response paging
}

// NewsResponse is created just for swagger, it's not uses in real (we just proxying requests/responses)
type NewsResponse struct {
	ID              uuid.UUID     `json:"id"`
	Title           string        `json:"title"`
	Content         string        `json:"content"`
	PublicationTime time.Time     `json:"publication_time"`
	Language        string        `json:"language"`
	Copyright       string        `json:"copyright"`
	Distributor     string        `json:"distributor"`
	SourceName      string        `json:"source_name"`
	Subjects        []string      `json:"subjects"`
	CompanyInfo     []CompanyInfo `json:"company_info"`
}

type CompanyInfo struct {
	Symbol      string `json:"symbol"`
	CompanyName string `json:"company_name"`
}

type NewsByIDRequest struct {
	ID string `json:"id"`
}

type SymbolSubjectsRequest struct {
	Symbol string `json:"symbol"`
}

type SymbolSubjectsResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TaxonomyCode struct {
	Code     string         `json:"code"`
	Name     string         `json:"name"`
	SubCodes []TaxonomyCode `json:"sub_codes,omitempty"`
}
