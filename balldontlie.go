package balldontlie

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// DefaultHostname for the balldontlie API
	DefaultHostname string = "https://www.balldontlie.io"
	userAgent       string = "go-balldontlie"
	maxHTTPCode     int    = 299
)

// Client - Used to interact with balldontlie api
type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string

	Games   *GameService
	Teams   *TeamService
	Players *PlayerService
	Stats   *StatService
}

// Response - Includes the Response for debugging purposes
type Response struct {
	*http.Response
	Container interface{}
}

// PageOpts - pagination query param options
type PageOpts struct {
	Page    int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
}

// PaginatedResponse - pagination metadata in the response
type PaginatedResponse struct {
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
	NextPage    int `json:"next_page"`
	PerPage     int `json:"per_page"`
	TotalCount  int `json:"total_count"`
}

// NewClient - Constructor for starting a new Client
func NewClient(baseURL string) (*Client, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client: &http.Client{
			Timeout: time.Second * 30, // Max of 30 secs
		},
		BaseURL:   url,
		UserAgent: userAgent,
	}

	c.Games = &GameService{client: c}
	c.Teams = &TeamService{client: c}
	c.Players = &PlayerService{client: c}
	c.Stats = &StatService{client: c}

	return c, nil
}

// NewRequest - creates a new HTTP request which can then be performed by Do.
func (c *Client) NewRequest(method, path string) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	reqURL := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest(method, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func newResponse(r *http.Response) *Response {
	resp := &Response{Response: r}
	return resp
}

// Do - performs a given request that was built with NewRequest. Return the
// response object
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	if resp.StatusCode > maxHTTPCode {
		errMsg := resp.Status + " returned from balldontlie"
		if strings.Contains(resp.Header.Get("Content-Type"), "application/json;") {
			err = json.NewDecoder(resp.Body).Decode(new(interface{}))
			if err != nil {
				return response, err
			}
		} else if strings.Contains(resp.Header.Get("Content-Type"), "text/plain;") {
			errBuf := &bytes.Buffer{}
			bufio.NewReader(resp.Body).WriteTo(errBuf)
			errMsg += ": " + errBuf.String()
		} else {
			errMsg += ": " + fmt.Sprintf("Response with unexpected Content-Type - %s received", resp.Header.Get("Content-Type"))
		}
		return response, fmt.Errorf(errMsg)
	}

	if v != nil {
		response.Container = v
	}

	if strings.Contains(resp.Header.Get("Content-Type"), "application/json;") {
		err = json.NewDecoder(resp.Body).Decode(response.Container)
		if err != nil {
			return response, err
		}
	} else {
		return response, fmt.Errorf("Response with unexpected Content-Type - %s received", resp.Header.Get("Content-Type"))
	}

	return response, nil
}
