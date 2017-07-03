package bitbucket

var HTTPClient contextKeyHTTPClient

type contextKeyHTTPClient struct{}

type Client interface {
	Projects() ([]Project, error)
	Users() ([]User, error)
	Repository(path string) (Repository, error)
}

//
//func NewClient(ctx context.Context, endpoint string) (*Client, error) {
//	c := internal.ContextClient(ctx)
//	if c == nil {
//		c = http.DefaultClient
//	}
//	ctx = context.WithValue(ctx, HTTPClient, c)
//
//	return &Client{d}, nil
//}

//func (c *Client) Projects() ([]Project, error) {
//	return c.driver.Projects()
//}
