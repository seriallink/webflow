package webflow

type QueueResponse struct {
	Queued bool `json:"queued"`
}

type DeleteResponse struct {
	Deleted int `json:"deleted"`
}
