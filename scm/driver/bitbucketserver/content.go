package bitbucketserver

import (
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
	lines := []string{}
	for _, v := range out.Lines {
		lines = append(lines, v.Text)
	}
	data := strings.Join(lines, "\n")
	return &scm.Content{
		Path: path,
		Data: []byte(data),
	}, res, err
}

//Bitbucket Server api: https://docs.atlassian.com/bitbucket-server/rest/5.0.1/bitbucket-rest.html#idm45993793705776
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

func (s *contentService) Update(ctx context.Context, repo, path string, params *scm.ContentParams) (*scm.Response, error) {
	projectRepo := strings.Split(repo, "/")
	endpoint := fmt.Sprintf("/rest/api/1.0/projects/%s/repos/%s/browse/%s", projectRepo[0], projectRepo[1], path)
	headerParams := make(map[string]string)
	formParams := url.Values{}

	formParams.Add("message", params.Message)
	formParams.Add("content", string(params.Data))
	formParams.Add("branch", params.Branch)
	formParams.Add("sourceCommitId", params.Branch)

	res, err := s.client.doForm(ctx, "PUT", endpoint, headerParams, formParams, nil)
	return res, err
}

func (s *contentService) Delete(ctx context.Context, repo, path, ref string) (*scm.Response, error) {
	return nil, scm.ErrNotSupported
}

func (s *contentService) List(ctx context.Context, repo, path, ref string, opts scm.ListOptions) ([]*scm.ContentInfo, *scm.Response, error) {
	return nil, nil, scm.ErrNotSupported
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
