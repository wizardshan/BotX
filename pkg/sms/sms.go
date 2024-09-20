package sms

type Client struct {
}

func NewClient() *Client {
	client := new(Client)
	return client
}

func (c *Client) Send(mobile string) error {
	return nil
}
