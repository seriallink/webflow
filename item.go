package webflow

import (
	"fmt"
	"strconv"
)

type Item map[string]interface{}

type ResponseItems struct {
	Count  int    `json:"count"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Total  int    `json:"total"`
	Items  []Item `json:"items"`
}

type RequestItem struct {
	Fields *Item `json:"fields,omitempty"`
}

func (i *Item) Get(key string) interface{} {
	return (*i)[key]
}

func (i *Item) Id() string {
	return i.Get("_id").(string)
}

func (i *Item) Name() string {
	return i.Get("name").(string)
}

func (i *Item) Slug() string {
	return i.Get("slug").(string)
}

func (i *Item) Archived() bool {
	return i.Get("_archived").(bool)
}

func (i *Item) Draft() bool {
	return i.Get("_draft").(bool)
}

func (c *Client) ListItems(collectionId string, offset, limit int) (*ResponseItems, error) {
	resp := new(ResponseItems)
	params := Params{
		"offset": strconv.Itoa(offset),
		"limit":  strconv.Itoa(limit),
	}
	if err := c.Get("collections/"+collectionId+"/items/", params, nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetItem(collectionId, itemId string) (*Item, error) {
	resp := new(ResponseItems)
	if err := c.Get("collections/"+collectionId+"/items/"+itemId, nil, nil, &resp); err != nil {
		return nil, err
	}
	if len(resp.Items) > 0 {
		return &resp.Items[0], nil
	}
	return nil, nil
}

func (c *Client) CreateItem(collectionId string, live bool, fields *RequestItem) (*Item, error) {
	item := new(Item)
	if err := c.Post(fmt.Sprintf("collections/%s/items?live=%s", collectionId, strconv.FormatBool(live)), fields, nil, item); err != nil {
		return nil, err
	}
	return item, nil
}

func (c *Client) UpdateItem(collectionId, itemId string, live bool, fields *RequestItem) (*Item, error) {
	item := new(Item)
	if err := c.Put(fmt.Sprintf("collections/%s/items/%s?live=%s", collectionId, itemId, strconv.FormatBool(live)), fields, nil, item); err != nil {
		return nil, err
	}
	return item, nil
}

func (c *Client) PatchItem(collectionId, itemId string, live bool, fields *RequestItem) (*Item, error) {
	item := new(Item)
	if err := c.Patch(fmt.Sprintf("collections/%s/items/%s?live=%s", collectionId, itemId, strconv.FormatBool(live)), fields, nil, item); err != nil {
		return nil, err
	}
	return item, nil
}

func (c *Client) RemoveItem(collectionId, itemId string, live bool) (*DeleteResponse, error) {
	resp := new(DeleteResponse)
	if err := c.Delete(fmt.Sprintf("collections/%s/items/%s?live=%s", collectionId, itemId, strconv.FormatBool(live)), nil, nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
