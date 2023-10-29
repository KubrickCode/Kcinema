package tmdb

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/k0kubun/pp"
)

type ProxyClient struct {
	baseURL string
	client  *http.Client
	token   string
}

func NewProxyClient() *ProxyClient {
	baseURL :=  "https://api.themoviedb.org"
	token := "TMDB_TOKEN"

	return &ProxyClient{
		baseURL: baseURL,
		client:  &http.Client{},
		token:   token,
	}
}

func (c *ProxyClient) NewRequest() *RequestBuilder {
	rb := NewRequestBuilder(c.baseURL)
	rb.AddHeader("Authorization", fmt.Sprintf("Bearer %s", c.token))
	return rb
}

func (c *ProxyClient) Do(req *http.Request, out interface{}) (bodyBytes []byte, pageInfo *string, err error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("do client api failed: %v", err)
	}

	if pageInfo, err = c.extractNextPageInfo(resp); err != nil {
		return nil, nil, fmt.Errorf("extract next page info failed: %v", err)
	}

	defer resp.Body.Close()

	bodyBytes, _ = io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		pp.Println(string(bodyBytes))
		return nil, pageInfo, fmt.Errorf("received non-2xx response status: %v body: %s", resp.Status, string(bodyBytes))
	}

	decoder := json.NewDecoder(bufio.NewReader(bytes.NewBuffer(bodyBytes)))
	err = decoder.Decode(out)
	if err != nil {
		pp.Println(string(bodyBytes))
		return nil, pageInfo, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(out)
	if err != nil {
		pp.Println(string(bodyBytes))
		return nil, pageInfo, err
	}

	return bodyBytes, pageInfo, nil
}

func (c *ProxyClient) extractNextPageInfo(resp *http.Response) (pageInfo *string, err error) {
	linkHeader := resp.Header.Get("Link")
	if linkHeader != "" {
		re := regexp.MustCompile(`<(.*?)>; rel="next"`)
		matches := re.FindStringSubmatch(linkHeader)
		if len(matches) == 2 {
			nextURL := matches[1]
			pageInfo = extractPageInfoFromURL(nextURL)
		}
	}
	return pageInfo, nil
}

func extractPageInfoFromURL(url string) *string {
	parameters := strings.Split(url, "&")
	for _, param := range parameters {
		keyValue := strings.Split(param, "=")
		if len(keyValue) == 2 && keyValue[0] == "page_info" {
			return &keyValue[1]
		}
	}
	return nil
}