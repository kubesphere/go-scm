package bitbucketserver

import (
	"context"
	"github.com/drone/go-scm/scm"
)

type pullService struct {
	*issueService
}

func (s *pullService) Find(ctx context.Context, repo string, number int) (*scm.PullRequest, *scm.Response, error) {
	return nil,nil,scm.ErrNotSupported
}

func (s *pullService) List(ctx context.Context, repo string, opts scm.PullRequestListOptions) ([]*scm.PullRequest, *scm.Response, error) {
	return nil,nil,scm.ErrNotSupported
}

func (s *pullService) ListChanges(ctx context.Context, repo string, number int, opts scm.ListOptions) ([]*scm.Change, *scm.Response, error) {
	return nil,nil,scm.ErrNotSupported
}

func (s *pullService) Merge(ctx context.Context, repo string, number int) (*scm.Response, error) {
	return nil,scm.ErrNotSupported
}

func (s *pullService) Close(ctx context.Context, repo string, number int) (*scm.Response, error) {
	return nil, scm.ErrNotSupported
}