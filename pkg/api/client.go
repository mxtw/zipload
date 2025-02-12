package api

type Client struct {
	Token string
	Host  string
}

func NewClient(token string, host string) Client {
	client := Client{
		token,
		host,
	}
	return client
}
