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
	SeperateHeader      bool     `json:"seperate_header"            require:"false"  inpath:"true"`
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
	Promotions          []string `json:"promotions"                 require:"false"  inpath:"true"`
	Item                string   `json:"item"                       require:"false"  inpath:"true"`
	TwitterEnabled      bool     `json:"twitter_enabled"            require:"false"  inpath:"true"`
	Template            string   `json:"template"                   require:"false"  inpath:"true"`
	MetadataLevel       int      `json:"metadata_level"             require:"false"  inpath:"true"`
	FundamentalLevel    int      `json:"fundamental_level"          require:"false"  inpath:"true"`
	SeperateHeader      bool     `json:"seperate_header"            require:"false"  inpath:"true"`
	DocGroup            string   `json:"doc_group"                  require:"true"   inpath:"true"`
	ArticleCacheTtl     int      `json:"article_cache_ttl"          require:"false"  inpath:"true"`
	EnableHtml          bool     `json:"enable_html"                require:"false"  inpath:"true"    default:"true"`
	PreferredDataSource string   `json:"preferred_data_source"      require:"false"  inpath:"true"`
	UseCustomerAssets   bool     `json:"use_customer_assets"        require:"false"  inpath:"true"`
	TwitterMode         string   `json:"twitter_mode"               require:"false"  inpath:"true"`
	Length              int      `json:"length"                     require:"false"  inpath:"true"`
	ImagesEnabled       bool     `json:"images_enabled"             require:"false"  inpath:"true"`
}

type HWatchlistRequest struct {
	Asset               string   `json:"asset"                      require:"true"   inpath:"false"`
	Promotions          []string `json:"promotions"                 require:"false"  inpath:"true"`
	Item                string   `json:"item"                       require:"false"  inpath:"true"`
	TwitterEnabled      bool     `json:"twitter_enabled"            require:"false"  inpath:"true"`
	Template            string   `json:"template"                   require:"false"  inpath:"true"`
	MetadataLevel       int      `json:"metadata_level"             require:"false"  inpath:"true"`
	FundamentalLevel    int      `json:"fundamental_level"          require:"false"  inpath:"true"`
	SeperateHeader      bool     `json:"seperate_header"            require:"false"  inpath:"true"`
	DocGroup            string   `json:"doc_group"                  require:"true"   inpath:"true"`
	ArticleCacheTtl     int      `json:"article_cache_ttl"          require:"false"  inpath:"true"`
	EnableHtml          bool     `json:"enable_html"                require:"false"  inpath:"true"    default:"true"`
	PreferredDataSource string   `json:"preferred_data_source"      require:"false"  inpath:"true"`
	UseCustomerAssets   bool     `json:"use_customer_assets"        require:"false"  inpath:"true"`
	TwitterMode         string   `json:"twitter_mode"               require:"false"  inpath:"true"`
	Length              int      `json:"length"                     require:"false"  inpath:"true"`
	ImagesEnabled       bool     `json:"images_enabled"             require:"false"  inpath:"true"`
}
