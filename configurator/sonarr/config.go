package sonarr

type selectOption struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}

type field struct {
	Order         int            `json:"order"`
	Name          string         `json:"name"`
	Label         string         `json:"label"`
	Value         interface{}    `json:"value,omitempty"`
	Type          string         `json:"type"`
	Advanced      bool           `json:"advanced"`
	HelpText      string         `json:"helpText,omitempty"`
	SelectOptions []selectOption `json:"selectOptions,omitempty"`
}

type config struct {
	Enable             bool    `json:"enable"`
	Protocol           string  `json:"protocol"`
	Name               string  `json:"name"`
	Fields             []field `json:"fields"`
	ImplementationName string  `json:"implementationName"`
	Implementation     string  `json:"implementation"`
	ConfigContract     string  `json:"configContract"`
	InfoLink           string  `json:"infoLink"`
	ID                 int     `json:"id"`
}
