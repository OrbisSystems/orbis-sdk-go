package model

type HSDailySummaryRequest struct {
	Asset               string   `json:"asset"                      require:"true"   inpath:"false"`
	Promotions          []string `json:"promotions"                 require:"false"  inpath:"true"`
	Item                string   `json:"item"                       require:"false"  inpath:"true"`
	TwitterEnabled      bool     `json:"twitter_enabled"            require:"false"  inpath:"true"`
	Template            string   `json:"template"                   require:"false"  inpath:"true"`
	MetadataLevel       int      `json:"metadata_level"             require:"false"  inpath:"true"`
	FundamentalLevel    int      `json:"fundamental_level"          require:"false"  inpath:"true"`
	SeperateHeader      bool     `json:"seperate_header"            require:"false"  inpath:"true"`
	DocGroup            string   `json:"doc_group"                  require:"false"  inpath:"true"`
	ArticleCacheTtl     int      `json:"article_cache_ttl"          require:"false"  inpath:"true"`
	EnableHtml          bool     `json:"enable_html"                require:"false"  inpath:"true"    default:"true"`
	PreferredDataSource string   `json:"preferred_data_source"      require:"false"  inpath:"true"`
	UseCustomerAssets   bool     `json:"use_customer_assets"        require:"false"  inpath:"true"`
	TwitterMode         string   `json:"twitter_mode"               require:"false"  inpath:"true"`
	Length              int      `json:"length"                     require:"false"  inpath:"true"`
	ImagesEnabled       bool     `json:"images_enabled"             require:"false"  inpath:"true"`
}

type HSDailySummaryResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		Text     string `json:"text"`
		Content  string `json:"content"`
		Metadata struct {
			GenerationUtcDate string      `json:"generation_utc_date"`
			GenerationRunTime interface{} `json:"generation_run_time"`
			CustomerName      string      `json:"customer_name"`
			ContentType       string      `json:"content_type"`
			DocStructure      string      `json:"doc_structure"`
			DocStructureScore float64     `json:"doc_structure_score"`
			Header            string      `json:"header"`
			SnippetsDict      struct {
			} `json:"snippets_dict"`
			AssetQuote struct {
				Timestamp     string  `json:"timestamp"`
				Open          float64 `json:"open"`
				PreviousClose float64 `json:"previous_close"`
				Close         float64 `json:"close"`
				High          float64 `json:"high"`
				Low           float64 `json:"low"`
			} `json:"asset_quote"`
			TargetSymbolsQuotesList []struct {
				Symbol               string  `json:"symbol"`
				Name                 string  `json:"name"`
				Price                float64 `json:"price"`
				ChangesPercentage    float64 `json:"changesPercentage"`
				Change               float64 `json:"change"`
				DayLow               float64 `json:"dayLow"`
				DayHigh              float64 `json:"dayHigh"`
				YearHigh             float64 `json:"yearHigh"`
				YearLow              float64 `json:"yearLow"`
				MarketCap            float64 `json:"marketCap"`
				PriceAvg50           float64 `json:"priceAvg50"`
				PriceAvg200          float64 `json:"priceAvg200"`
				Exchange             string  `json:"exchange"`
				Volume               int     `json:"volume"`
				AvgVolume            int     `json:"avgVolume"`
				Open                 float64 `json:"open"`
				PreviousClose        float64 `json:"previousClose"`
				Eps                  float64 `json:"eps"`
				Pe                   float64 `json:"pe"`
				EarningsAnnouncement string  `json:"earningsAnnouncement"`
				SharesOutstanding    float64 `json:"sharesOutstanding"`
				Date                 string  `json:"date"`
				AssetClass           string  `json:"asset_class"`
			} `json:"target_symbols_quotes_list"`
		} `json:"metadata"`
	} `json:"result"`
}
