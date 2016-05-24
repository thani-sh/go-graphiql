package graphiql

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// GraphQLError ...
type GraphQLError struct {
	Message string `json:"message"`
}

// GraphQLError implements error interface
func (e *GraphQLError) Error() string {
	return e.Message
}

// Request ...
type Request struct {
	Query string `json:"query"`
}

// Response ...
type Response struct {
	Data   *json.RawMessage `json:"data"`
	Errors []GraphQLError   `json:"errors,omitempty"`
}

// Client ...
type Client struct {
	Endpoint string
	Header   http.Header
	Client   http.Client
}

// NewClient ...
func NewClient(uri string) (c *Client, err error) {
	c = &Client{
		Endpoint: uri,
		Header:   http.Header{},
		Client:   http.Client{},
	}
	c.Header.Set("Content-Type", "application/json")

	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	if u.User == nil {
		return c, nil
	}

	user := u.User.Username()
	pass, hasPass := u.User.Password()
	if user != "" && hasPass && pass != "" {
		req := http.Request{Header: c.Header}
		req.SetBasicAuth(user, pass)
	}

	u.User = nil
	c.Endpoint = u.String()

	return c, nil
}

// Send ...
func (c *Client) Send(req *Request) (res *Response, err error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	bodyBuf := bytes.NewBuffer(body)
	httpReq, err := http.NewRequest(http.MethodPost, c.Endpoint, bodyBuf)
	if err != nil {
		return nil, err
	}
	httpReq.Header = c.Header

	httpRes, err := c.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	res = &Response{}
	if err := json.NewDecoder(httpRes.Body).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}
