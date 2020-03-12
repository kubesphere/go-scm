// Package bitbucket implements a Bitbucket Server client.
package bitbucketserver

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/drone/go-scm/scm"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// New returns a new Bitbucket Server API client.
func New(uri string) (*scm.Client, error) {
	base, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(base.Path, "/") {
		base.Path = base.Path + "/"
	}
	client := &wrapper{new(scm.Client)}
	client.BaseURL = base
	// initialize services
	client.Driver = scm.DriverBitbucketServer
	client.Linker = &linker{base.String()}
	client.Contents = &contentService{client}
	client.Git = &gitService{client}
	client.Issues = &issueService{client}
	client.Organizations = &organizationService{client}
	client.PullRequests = &pullService{&issueService{client}}
	client.Repositories = &repositoryService{client}
	client.Reviews = &reviewService{client}
	client.Users = &userService{client}
	client.Webhooks = &webhookService{client}
	return client.Client, nil
}

// NewDefault returns a new Bitbucket Server API client just for unit test
func NewDefault() *scm.Client {
	client, _ := New("https://localhost.api.bitbucket.server")
	return client
}

// wraper wraps the Client to provide high level helper functions
// for making http requests and unmarshaling the response.
type wrapper struct {
	*scm.Client
}

func (c *wrapper) doForm(ctx context.Context, method, path string, headerParams map[string]string,
	formParams url.Values, out interface{}) (*scm.Response, error) {

	req := &scm.Request{
		Method:   method,
		Path:     path,
		PostForm: formParams,
		Header:   http.Header{},
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		for h, v := range headerParams {
			req.Header.Set(h, v)
		}
	}

	return c.callApi(ctx, req, out)

}

// do wraps the Client.Do function by creating the Request and
// unmarshalling the response.
func (c *wrapper) do(ctx context.Context, method, path string, in, out interface{}) (*scm.Response, error) {
	req := &scm.Request{
		Method: method,
		Path:   path,
	}
	// if we are posting or putting data, we need to
	// write it to the body of the request.
	if in != nil {
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(in)
		req.Header = map[string][]string{
			"Content-Type": {"application/json"},
		}
		req.Body = buf
	}

	return c.callApi(ctx, req, out)
}

func (c *wrapper) callApi(ctx context.Context, req *scm.Request, out interface{}) (*scm.Response, error) {
	// execute the http request
	res, err := c.Client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// if an error is encountered, unmarshal and return the
	// error response.
	if res.Status == 401 {
		return res, scm.ErrNotAuthorized
	} else if res.Status > 300 {
		err := new(Error)
		json.NewDecoder(res.Body).Decode(err)
		return res, err
	}

	if out == nil {
		return res, nil
	}

	// if raw output is expected, copy to the provided
	// buffer and exit.
	if w, ok := out.(io.Writer); ok {
		io.Copy(w, res.Body)
		return res, nil
	}

	// if a json response is expected, parse and return
	// the json response.
	return res, json.NewDecoder(res.Body).Decode(out)
}

// Error represents a Bitbucket Server error.
type Error struct {
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
