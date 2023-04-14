package webflow

type Application struct {
	Id          string `json:"_id,omitempty"`
	Description string `json:"description,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
	Name        string `json:"name,omitempty"`
	Owner       string `json:"owner,omitempty"`
	OwnerType   string `json:"ownerType,omitempty"`
}
