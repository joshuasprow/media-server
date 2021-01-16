package indexer

// Args ...
type Args struct {
	URL    string
	APIKey string
}

// Field ...
type Field struct {
	Order    int         `json:"order"`
	Name     string      `json:"name"`
	Label    string      `json:"label"`
	Value    interface{} `json:"value,omitempty"`
	Type     string      `json:"type"`
	Advanced bool        `json:"advanced"`
	HelpText string      `json:"helpText,omitempty"`
	Unit     string      `json:"unit,omitempty"`
}

// Config ...
type Config struct {
	EnableRss          bool    `json:"enableRss"`
	EnableSearch       bool    `json:"enableSearch"`
	SupportsRss        bool    `json:"supportsRss"`
	SupportsSearch     bool    `json:"supportsSearch"`
	Protocol           string  `json:"protocol"`
	Name               string  `json:"name"`
	Fields             []Field `json:"fields"`
	ImplementationName string  `json:"implementationName"`
	Implementation     string  `json:"implementation"`
	ConfigContract     string  `json:"configContract"`
	InfoLink           string  `json:"infoLink"`
	ID                 int     `json:"id"`
}

func newConfig(args Args) Config {
	urlField := Field{
		Order:    0,
		Name:     "BaseUrl",
		Label:    "URL",
		Value:    args.URL,
		Type:     "textbox",
		Advanced: false,
	}

	apiKeyField := Field{
		Order:    2,
		Name:     "ApiKey",
		Label:    "API Key",
		Value:    args.APIKey,
		Type:     "textbox",
		Advanced: false,
	}

	fields := []Field{
		urlField,
		{
			Order:    1,
			Name:     "ApiPath",
			Label:    "API Path",
			HelpText: "Path to the api, usually /api",
			Value:    "/api",
			Type:     "textbox",
			Advanced: true,
		},
		apiKeyField,
		{
			Order:    3,
			Name:     "Categories",
			Label:    "Categories",
			HelpText: "Comma Separated list, leave blank to disable standard/daily shows",
			Value:    []int{5000, 5010, 5040},
			Type:     "textbox",
			Advanced: false,
		},
		{
			Order:    4,
			Name:     "AnimeCategories",
			Label:    "Anime Categories",
			HelpText: "Comma Separated list, leave blank to disable anime",
			Value:    []string{},
			Type:     "textbox",
			Advanced: false,
		},
		{
			Order:    5,
			Name:     "AdditionalParameters",
			Label:    "Additional Parameters",
			HelpText: "Additional Newznab parameters",
			Type:     "textbox",
			Advanced: true,
		},
		{
			Order:    6,
			Name:     "MinimumSeeders",
			Label:    "Minimum Seeders",
			HelpText: "Minimum number of seeders required.",
			Value:    1,
			Type:     "textbox",
			Advanced: true,
		},
		{
			Order:    7,
			Name:     "SeedCriteria.SeedRatio",
			Label:    "Seed Ratio",
			HelpText: "The ratio a torrent should reach before stopping, empty is download client's default",
			Type:     "textbox",
			Advanced: true,
		},
		{
			Order:    8,
			Name:     "SeedCriteria.SeedTime",
			Label:    "Seed Time",
			Unit:     "minutes",
			HelpText: "The time a torrent should be seeded before stopping, empty is download client's default",
			Type:     "textbox",
			Advanced: true,
		},
		{
			Order:    9,
			Name:     "SeedCriteria.SeasonPackSeedTime",
			Label:    "Season-Pack Seed Time",
			Unit:     "minutes",
			HelpText: "The time a torrent should be seeded before stopping, empty is download client's default",
			Type:     "textbox",
			Advanced: true,
		},
	}

	return Config{
		EnableRss:          true,
		EnableSearch:       true,
		SupportsRss:        true,
		SupportsSearch:     true,
		Protocol:           "torrent",
		Name:               "Jackett 2",
		Fields:             fields,
		ImplementationName: "Torznab",
		Implementation:     "Torznab",
		ConfigContract:     "TorznabSettings",
		InfoLink:           "https://github.com/Sonarr/Sonarr/wiki/Supported-Indexers#torznab",
		ID:                 0,
	}
}
