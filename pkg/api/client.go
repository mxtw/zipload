package api

type Client struct {
	token string
	host  string
}

func NewClient(token string, host string) Client {
	client := Client{
		token,
		host,
	}
	return client
}
