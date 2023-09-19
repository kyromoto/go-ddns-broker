package core

type Client struct {
	Entity
	description string
	password    string
}

func (c *Client) AssertPassword(password string) (ok bool) {
	return c.password == password
}
