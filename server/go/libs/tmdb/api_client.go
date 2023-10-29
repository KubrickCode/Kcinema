package tmdb

const APIVersion = "2023-10"

type APIClient struct {
	c *ProxyClient
}

func NewAPIClient() *APIClient {
	return &APIClient{
		c: NewProxyClient(),
	}
}

func (c *APIClient) GetMovieList() (movies []MovieList, err error) {
	return GetMovieList(c.c)
}