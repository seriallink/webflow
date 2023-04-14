package webflow

import "time"

type Collection struct {
	Id           string    `json:"_id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Slug         string    `json:"slug,omitempty"`
	SingularName string    `json:"singularName,omitempty"`
	LastUpdated  time.Time `json:"lastUpdated,omitempty"`
	CreatedOn    time.Time `json:"createdOn,omitempty"`
	Fields       []Field   `json:"fields,omitempty"`
}

func (c *Client) ListCollections(siteId string) (collections []Collection, err error) {
	err = c.Get("sites/"+siteId+"/collections/", nil, nil, &collections)
	return
}

func (c *Client) GetCollection(collectionId string) (*Collection, error) {
	collection := new(Collection)
	if err := c.Get("collections/"+collectionId, nil, nil, collection); err != nil {
		return nil, err
	}
	return collection, nil
}
