package downloadclient

// Args ...
type Args struct {
	Username string
	Password string
}

// SelectOption ...
type SelectOption struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}

// Field ...
type Field struct {
	Order         int            `json:"order"`
	Name          string         `json:"name"`
	Label         string         `json:"label"`
	Value         interface{}    `json:"value,omitempty"`
	Type          string         `json:"type"`
	Advanced      bool           `json:"advanced"`
	HelpText      string         `json:"helpText,omitempty"`
	SelectOptions []SelectOption `json:"selectOptions,omitempty"`
}

// Config ...
type Config struct {
	Enable             bool    `json:"enable"`
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
	userField := Field{
		Order:    3,
		Name:     "Username",
		Label:    "Username",
		Value:    args.Username,
		Type:     "textbox",
		Advanced: false,
	}
	passField := Field{
		Order:    4,
		Name:     "Password",
		Label:    "Password",
		Value:    args.Password,
		Type:     "password",
		Advanced: false,
	}

	fields := []Field{
		{
			Order:    0,
			Name:     "Host",
			Label:    "Host",
			Value:    "transmission",
			Type:     "textbox",
			Advanced: false,
		},
		{
			Order:    1,
			Name:     "Port",
			Label:    "Port",
			Value:    9091,
			Type:     "textbox",
			Advanced: false,
		},
		{
			Order:    2,
			Name:     "UrlBase",
			Label:    "Url Base",
			HelpText: "Adds a prefix to the transmission rpc url, eg http://[host]:[port]/[urlBase]/rpc, defaults to '/transmission/'",
			Value:    "/transmission/",
			Type:     "textbox",
			Advanced: true,
		},
		userField,
		passField,
		{
			Order:    5,
			Name:     "TvCategory",
			Label:    "Category",
			HelpText: "Adding a category specific to Sonarr avoids conflicts with unrelated downloads, but it's optional. Creates a [category] subdirectory in the output directory.",
			Type:     "textbox",
			Advanced: false,
		},
		{
			Order:    6,
			Name:     "TvDirectory",
			Label:    "Directory",
			HelpText: "Optional location to put downloads in, leave blank to use the default Transmission location",
			Type:     "textbox",
			Advanced: true,
		},
		{
			Order:    7,
			Name:     "RecentTvPriority",
			Label:    "Recent Priority",
			HelpText: "Priority to use when grabbing episodes that aired within the last 14 days",
			Value:    0,
			Type:     "select",
			Advanced: false,
			SelectOptions: []SelectOption{
				{Value: 0, Name: "Last"},
				{Value: 1, Name: "First"},
			},
		},
		{
			Order:    8,
			Name:     "OlderTvPriority",
			Label:    "Older Priority",
			HelpText: "Priority to use when grabbing episodes that aired over 14 days ago",
			Value:    0,
			Type:     "select",
			Advanced: false,
			SelectOptions: []SelectOption{
				{Value: 0, Name: "Last"},
				{Value: 1, Name: "First"},
			},
		},
		{
			Order:    9,
			Name:     "AddPaused",
			Label:    "Add Paused",
			Value:    false,
			Type:     "checkbox",
			Advanced: false,
		},
		{
			Order:    10,
			Name:     "UseSsl",
			Label:    "Use SSL",
			Value:    false,
			Type:     "checkbox",
			Advanced: false,
		},
	}

	return Config{
		Enable:             true,
		Protocol:           "torrent",
		Name:               "Transmission",
		Fields:             fields,
		ImplementationName: "Transmission",
		Implementation:     "Transmission",
		ConfigContract:     "TransmissionSettings",
		InfoLink:           "https://github.com/Sonarr/Sonarr/wiki/Supported-s#transmission",
		ID:                 0,
	}
}
