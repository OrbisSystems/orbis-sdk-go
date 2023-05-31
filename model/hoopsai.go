package model

type HSDailySummaryRequest struct {
	Asset               string   `json:"asset"                      require:"true"   inpath:"false"`
	Promotions          []string `json:"promotions"                 require:"false"  inpath:"true"`
	Item                string   `json:"item"                       require:"false"  inpath:"true"`
	TwitterEnabled      bool     `json:"twitter_enabled"            require:"false"  inpath:"true"`
	Template            string   `json:"template"                   require:"false"  inpath:"true"`
	MetadataLevel       int      `json:"metadata_level"             require:"false"  inpath:"true"`
	FundamentalLevel    int      `json:"fundamental_level"          require:"false"  inpath:"true"`
	SeparateHeader      bool     `json:"separate_header"            require:"false"  inpath:"true"`
	DocGroup            string   `json:"doc_group"                  require:"true"   inpath:"true"`
	ArticleCacheTtl     int      `json:"article_cache_ttl"          require:"false"  inpath:"true"`
	EnableHtml          bool     `json:"enable_html"                require:"false"  inpath:"true"    default:"true"`
	PreferredDataSource string   `json:"preferred_data_source"      require:"false"  inpath:"true"`
	UseCustomerAssets   bool     `json:"use_customer_assets"        require:"false"  inpath:"true"`
	TwitterMode         string   `json:"twitter_mode"               require:"false"  inpath:"true"`
	Length              int      `json:"length"                     require:"false"  inpath:"true"`
	ImagesEnabled       bool     `json:"images_enabled"             require:"false"  inpath:"true"`
}

type HSWeeklySummaryRequest struct {
	Asset               string   `json:"asset"                      require:"true"   inpath:"false"`
	Promotions          []string `json:"promotions"                 require:"false"  inpath:"true"`
	Item                string   `json:"item"                       require:"false"  inpath:"true"`
	TwitterEnabled      bool     `json:"twitter_enabled"            require:"false"  inpath:"true"`
	Template            string   `json:"template"                   require:"false"  inpath:"true"`
	MetadataLevel       int      `json:"metadata_level"             require:"false"  inpath:"true"`
	FundamentalLevel    int      `json:"fundamental_level"          require:"false"  inpath:"true"`
	SeparateHeader      bool     `json:"separate_header"            require:"false"  inpath:"true"`
	DocGroup            string   `json:"doc_group"                  require:"true"   inpath:"true"`
	ArticleCacheTtl     int      `json:"article_cache_ttl"          require:"false"  inpath:"true"`
	EnableHtml          bool     `json:"enable_html"                require:"false"  inpath:"true"    default:"true"`
	PreferredDataSource string   `json:"preferred_data_source"      require:"false"  inpath:"true"`
	UseCustomerAssets   bool     `json:"use_customer_assets"        require:"false"  inpath:"true"`
	TwitterMode         string   `json:"twitter_mode"               require:"false"  inpath:"true"`
	Length              int      `json:"length"                     require:"false"  inpath:"true"`
	ImagesEnabled       bool     `json:"images_enabled"             require:"false"  inpath:"true"`
}

type HSPortfolioRequest struct {
	Promotions          []string `json:"promotions"`
	Item                string   `json:"item"`
	TwitterEnabled      bool     `json:"twitter_enabled"`
	Template            string   `json:"template"`
	MetadataLevel       int      `json:"metadata_level"`
	FundamentalLevel    int      `json:"fundamental_level"`
	SeparateHeader      bool     `json:"separate_header"`
	DocGroup            string   `json:"doc_group"`
	ArticleCacheTtl     int      `json:"article_cache_ttl"`
	EnableHtml          bool     `json:"enable_html"`
	PreferredDataSource string   `json:"preferred_data_source"`
	UseCustomerAssets   bool     `json:"use_customer_assets"`
	TwitterMode         string   `json:"twitter_mode"`
	Length              int      `json:"length"`
	ImagesEnabled       bool     `json:"images_enabled"`
}

type HSWatchlistRequest struct {
	Asset               string   `json:"asset"                      require:"true"   inpath:"false"`
	Promotions          []string `json:"promotions"                 require:"false"  inpath:"true"`
	Item                string   `json:"item"                       require:"false"  inpath:"true"`
	TwitterEnabled      bool     `json:"twitter_enabled"            require:"false"  inpath:"true"`
	Template            string   `json:"template"                   require:"false"  inpath:"true"`
	MetadataLevel       int      `json:"metadata_level"             require:"false"  inpath:"true"`
	FundamentalLevel    int      `json:"fundamental_level"          require:"false"  inpath:"true"`
	SeparateHeader      bool     `json:"separate_header"            require:"false"  inpath:"true"`
	DocGroup            string   `json:"doc_group"                  require:"true"   inpath:"true"`
	ArticleCacheTtl     int      `json:"article_cache_ttl"          require:"false"  inpath:"true"`
	EnableHtml          bool     `json:"enable_html"                require:"false"  inpath:"true"    default:"true"`
	PreferredDataSource string   `json:"preferred_data_source"      require:"false"  inpath:"true"`
	UseCustomerAssets   bool     `json:"use_customer_assets"        require:"false"  inpath:"true"`
	TwitterMode         string   `json:"twitter_mode"               require:"false"  inpath:"true"`
	Length              int      `json:"length"                     require:"false"  inpath:"true"`
	ImagesEnabled       bool     `json:"images_enabled"             require:"false"  inpath:"true"`
}

type HSWatchlistByUserAndNameRequest struct {
	UserID              string   `json:"user_id"                    require:"true"   inpath:"false"`
	WatchlistName       string   `json:"watchlist_name"             require:"true"   inpath:"false"`
	Promotions          []string `json:"promotions"                 require:"false"  inpath:"true"`
	Item                string   `json:"item"                       require:"false"  inpath:"true"`
	TwitterEnabled      bool     `json:"twitter_enabled"            require:"false"  inpath:"true"`
	Template            string   `json:"template"                   require:"false"  inpath:"true"`
	MetadataLevel       int      `json:"metadata_level"             require:"false"  inpath:"true"`
	FundamentalLevel    int      `json:"fundamental_level"          require:"false"  inpath:"true"`
	SeparateHeader      bool     `json:"separate_header"            require:"false"  inpath:"true"`
	DocGroup            string   `json:"doc_group"                  require:"true"   inpath:"true"`
	ArticleCacheTtl     int      `json:"article_cache_ttl"          require:"false"  inpath:"true"`
	EnableHtml          bool     `json:"enable_html"                require:"false"  inpath:"true"    default:"true"`
	PreferredDataSource string   `json:"preferred_data_source"      require:"false"  inpath:"true"`
	UseCustomerAssets   bool     `json:"use_customer_assets"        require:"false"  inpath:"true"`
	TwitterMode         string   `json:"twitter_mode"               require:"false"  inpath:"true"`
	Length              int      `json:"length"                     require:"false"  inpath:"true"`
	ImagesEnabled       bool     `json:"images_enabled"             require:"false"  inpath:"true"`
}

type HSMarketResearchRequest struct {
	Filters                string   `json:"filters"                       require:"false"  inpath:"true"`
	UseFullScreenerResults bool     `json:"use_full_screener_results"     require:"false"  inpath:"true"`
	Promotions             []string `json:"promotions"                    require:"false"  inpath:"true"`
	Item                   string   `json:"item"                          require:"false"  inpath:"true"`
	TwitterEnabled         bool     `json:"twitter_enabled"               require:"false"  inpath:"true"`
	Template               string   `json:"template"                      require:"false"  inpath:"true"`
	MetadataLevel          int      `json:"metadata_level"                require:"false"  inpath:"true"`
	FundamentalLevel       int      `json:"fundamental_level"             require:"false"  inpath:"true"`
	SeparateHeader         bool     `json:"separate_header"               require:"false"  inpath:"true"`
	DocGroup               string   `json:"doc_group"                     require:"true"   inpath:"true"`
	ArticleCacheTtl        int      `json:"article_cache_ttl"             require:"false"  inpath:"true"`
	EnableHtml             bool     `json:"enable_html"                   require:"false"  inpath:"true"    default:"true"`
	PreferredDataSource    string   `json:"preferred_data_source"         require:"false"  inpath:"true"`
	UseCustomerAssets      bool     `json:"use_customer_assets"           require:"false"  inpath:"true"`
	TwitterMode            string   `json:"twitter_mode"                  require:"false"  inpath:"true"`
	Length                 int      `json:"length"                        require:"false"  inpath:"true"`
	ImagesEnabled          bool     `json:"images_enabled"                require:"false"  inpath:"true"`
}
