package webflow

type User struct {
	Id        string `json:"_id,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

func (c *Client) AuthorizedUser() (*User, error) {
	resp := struct {
		User *User `json:"user"`
	}{}
	if err := c.Get("user/", nil, nil, &resp); err != nil {
		return nil, err
	}
	return resp.User, nil
}
