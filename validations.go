package webflow

type Validations struct {
	CollectionId  string `json:"collectionId,omitempty"`
	Format        string `json:"format,omitempty"`
	Precision     int    `json:"precision,omitempty"`
	MaxLength     int    `json:"maxLength,omitempty"`
	AllowNegative bool   `json:"allowNegative,omitempty"`
	SingleLine    bool   `json:"singleLine,omitempty"`
}
