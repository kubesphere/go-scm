package bitbucketserver

import (
	"bytes"
	"context"
	"fmt"
	"github.com/drone/go-scm/scm"
	"net/url"
	"strings"
)

type contentService struct {
	client *wrapper
}

func (s *contentService) Find(ctx context.Context, repo, path, ref string) (*scm.Content, *scm.Response, error) {
	projectRepo := strings.Split(repo, "/")
	endpoint := fmt.Sprintf("/rest/api/1.0/projects/%s/repos/%s/browse/%s?at=%s", projectRepo[0], projectRepo[1], path, ref)
	out := new(content)
	res, err := s.client.do(ctx, "GET", endpoint, nil, out)
	data := stringBuffer(out.Lines)
	return &scm.Content{
		Path: path,
		Data: []byte(data),
	}, res, err
}

func (s *contentService) Create(ctx context.Context, repo, path string, params *scm.ContentParams) (*scm.Response, error) {
	projectRepo := strings.Split(repo, "/")
	endpoint := fmt.Sprintf("/rest/api/1.0/projects/%s/repos/%s/browse/%s", projectRepo[0], projectRepo[1], path)
	headerParams := make(map[string]string)
	formParams := url.Values{}

	formParams.Add("message", params.Message)
	formParams.Add("content", string(params.Data))
	formParams.Add("branch", params.Branch)

	res, err := s.client.doForm(ctx, "PUT", endpoint, headerParams, formParams, nil)
	return res, err
}

// TODO
func (s *contentService) Update(ctx context.Context, repo, path string, params *scm.ContentParams) (*scm.Response, error) {
	//projectRepo := strings.Split(repo, "/")
	//endpoint := fmt.Sprintf("/rest/api/1.0/projects/%s/repos/%s/browse/%s?at=%s", projectRepo[0], projectRepo[1], path, params.Branch)
	return nil, scm.ErrNotSupported
}

func (s *contentService) Delete(ctx context.Context, repo, path, ref string) (*scm.Response, error) {
	return nil, scm.ErrNotSupported
}

func (s *contentService) List(ctx context.Context, repo, path, ref string, opts scm.ListOptions) ([]*scm.ContentInfo, *scm.Response, error) {
	return nil, nil, scm.ErrNotSupported
}

func stringBuffer(lines []contentLines) string {
	var b bytes.Buffer
	length := len(lines)
	for i, line := range lines {
		b.WriteString(line.Text)
		if i != length-1 {
			b.WriteString("\n")
		}

	}
	return b.String()
}

type content struct {
	Lines      []contentLines `json:"lines"`
	Start      int            `json:"start"`
	Size       int            `json:"size"`
	IsLastPage bool           `json:"isLastPage"`
}

type contentLines struct {
	Text string `json:"text"`
}
