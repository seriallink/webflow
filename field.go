package webflow

type Field struct {
	Id          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Slug        string       `json:"slug,omitempty"`
	Type        string       `json:"type,omitempty"`
	HelpText    string       `json:"helpText,omitempty"`
	Editable    bool         `json:"editable,omitempty"`
	Required    bool         `json:"required,omitempty"`
	Validations *Validations `json:"validations,omitempty"`
}
