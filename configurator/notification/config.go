package notification

// Args ...
type Args struct {
	Username string
	Password string
}

// Field ...
type Field struct {
	Order    int         `json:"order"`
	Name     string      `json:"name"`
	Label    string      `json:"label"`
	Value    interface{} `json:"value"` // could be any JSON type
	Type     string      `json:"type"`
	Advanced bool        `json:"advanced"`
	HelpText string      `json:"helpText,omitempty"`
}

// Config ...
type Config struct {
	OnGrab             bool          `json:"onGrab"`
	OnDownload         bool          `json:"onDownload"`
	OnUpgrade          bool          `json:"onUpgrade"`
	OnRename           bool          `json:"onRename"`
	SupportsOnGrab     bool          `json:"supportsOnGrab"`
	SupportsOnDownload bool          `json:"supportsOnDownload"`
	SupportsOnUpgrade  bool          `json:"supportsOnUpgrade"`
	SupportsOnRename   bool          `json:"supportsOnRename"`
	Tags               []interface{} `json:"tags"` // haven't used them yet
	Name               string        `json:"name"`
	Fields             []Field       `json:"fields"`
	ImplementationName string        `json:"implementationName"`
	Implementation     string        `json:"implementation"`
	ConfigContract     string        `json:"configContract"`
	InfoLink           string        `json:"infoLink"`
	ID                 int           `json:"id"`
}

func newConfig(args Args) Config {
	userField := Field{
		Order:    2,
		Name:     "Username",
		Label:    "Username",
		Value:    args.Username,
		Type:     "textbox",
		Advanced: false,
	}
	passField := Field{
		Order:    3,
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
			Value:    "plex",
			Type:     "textbox",
			Advanced: false,
		},
		{
			Order:    1,
			Name:     "Port",
			Label:    "Port",
			Value:    32400,
			Type:     "textbox",
			Advanced: false,
		},
		userField,
		passField,
		{
			Order:    4,
			Name:     "UpdateLibrary",
			Label:    "Update Library",
			Value:    true,
			Type:     "checkbox",
			Advanced: false,
		},
		{
			Order:    5,
			Name:     "UseSsl",
			Label:    "Use SSL",
			HelpText: "Connect to Plex over HTTPS instead of HTTP",
			Value:    false,
			Type:     "checkbox",
			Advanced: false,
		},
	}

	return Config{
		OnGrab:             false,
		OnDownload:         true,
		OnUpgrade:          true,
		OnRename:           true,
		SupportsOnGrab:     false,
		SupportsOnDownload: true,
		SupportsOnUpgrade:  true,
		SupportsOnRename:   true,
		Name:               "Plex",
		ImplementationName: "Plex Media Server",
		Implementation:     "PlexServer",
		ConfigContract:     "PlexServerSettings",
		InfoLink:           "https://github.com/Sonarr/Sonarr/wiki/Supported-Notifications#plexserver",
		ID:                 0,
		Fields:             fields,
		Tags:               []interface{}{},
	}
}
